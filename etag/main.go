package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	addr := ":12345"
	engine := gin.New()
	dir := "./static"

	cfg := ETagConfig{
		Dir:    dir,
		Prefix: "/static",
	}
	//staticRouter := engine.Group("/static").Use(NewETagHandler(cfg))
	//staticRouter.StaticFS("/", http.Dir(dir))
	engine.Use(NewETagHandler(cfg)).StaticFS("/static", http.Dir(dir))

	fmt.Println("http server listening on " + addr)
	engine.Run(addr)
}

type ETagConfig struct {
	Dir      string
	Prefix   string
	Index    string
	Strategy string
	MagAge   int
}

type etag struct {
	locker sync.RWMutex
	kv     map[string]string

	prefix             string
	cacheControlHeader string
}

func NewETagHandler(c ETagConfig) gin.HandlerFunc {
	e, err := newETag(c)
	if err != nil {
		panic("etag handler: " + err.Error())
	}

	return func(ctx *gin.Context) {
		e.handle(ctx)
	}
}

func newETag(c ETagConfig) (*etag, error) {
	e := &etag{
		kv:     make(map[string]string),
		prefix: c.Prefix,
	}
	if c.Dir == "" {
		return e, nil
	}
	dir := strings.Replace(c.Dir, "./", "", 1)
	var index string
	if c.Index == "" {
		index = "/index.html"
	}
	if c.Strategy == "" {
		e.cacheControlHeader = "no-cache"
	} else {
		e.cacheControlHeader = c.Strategy
	}
	if c.MagAge > 0 {
		e.cacheControlHeader += fmt.Sprintf(", max-age=%d", c.MagAge)
	}

	err := filepath.Walk(c.Dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			buffer, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			v := e.genHash(buffer)
			k := strings.Replace(path, dir, e.prefix, 1)
			e.kv[k] = v
			if strings.HasSuffix(k, index) {
				k := strings.TrimRight(k, index)
				e.kv[k] = v
				e.kv[k+"/"] = v
			}

			return nil
		})
	if err != nil {
		return e, err
	}

	return e, nil
}

func (e *etag) get(k string) string {
	e.locker.RLock()
	defer e.locker.RUnlock()

	v, _ := e.kv[k]
	return v
}

func (e *etag) genHash(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func (e *etag) handle(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	if !strings.HasPrefix(path, e.prefix) {
		ctx.Next()
		return
	}

	tag := e.get(path)
	matchTag := ctx.Request.Header.Get("if-none-match")
	ctx.Writer.Header().Set("etag", tag)
	ctx.Writer.Header().Set("Cache-Control", "no-cache, max-age=604800")

	if matchTag == tag {
		ctx.Status(http.StatusNotModified)
		ctx.Abort()
		return
	}

	ctx.Next()
}
