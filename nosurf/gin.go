package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/nosurf"
	adapter "github.com/gwatts/gin-adapter"
)

var templateString string = `
<!doctype html>
<html>
<body>
{{ if .name }}
<p>Your name: {{ .name }}</p>
{{ end }}
<form action="/" method="POST">
<input type="text" name="name">

<!-- Try removing this or changing its value
     and see what happens -->
<input type="hidden" name="csrf_token" value="{{ .token }}">
<input type="submit" value="Send">
</form>
</body>
</html>
`
var templ = template.Must(template.New("t1").Parse(templateString))

func index(c *gin.Context) {
	context := make(map[string]string)
	context["token"] = nosurf.Token(c.Request)

	templ.Execute(c.Writer, context)
}

func postIndex(c *gin.Context) {
	context := make(map[string]string)
	context["token"] = nosurf.Token(c.Request)
	context["name"] = c.PostForm("name")

	templ.Execute(c.Writer, context)
}

func main() {
	engine := gin.Default()

	MaxAge := 365 * 24 * 60 * 60

	baseCookie := http.Cookie{}
	baseCookie.MaxAge = MaxAge
	baseCookie.Name = "x_test"
	csrfHandler := CSRF(Config{
		BaseCookie: &baseCookie,
	})
	engine.Use(csrfHandler)
	engine.GET("/", index)
	engine.POST("/", postIndex)

	fmt.Println("Listening on http://127.0.0.1:8000/")
	engine.Run(":8000")
}

func CSRF(cfg Config) gin.HandlerFunc {
	next, a := adapter.New()

	h := nosurf.New(next)
	if cfg.FailureHandler != nil {
		h.SetFailureHandler(cfg.FailureHandler)
	}
	if cfg.BaseCookie != nil {
		h.SetBaseCookie(*cfg.BaseCookie)
	}
	for _, ig := range cfg.IgnorePaths {
		h.IgnorePath(ig)
	}
	h.IgnoreGlobs(cfg.IgnoreGlobs...)
	for _, dg := range cfg.DisableGlobs {
		h.DisablePath(dg)
	}
	h.DisableGlobs(cfg.DisableGlobs...)

	return a(h)
}

type Config struct {
	FailureHandler http.Handler
	IgnorePaths    []string
	IgnoreGlobs    []string
	DisablePaths   []string
	DisableGlobs   []string

	BaseCookie *http.Cookie
}
