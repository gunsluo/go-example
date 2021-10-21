package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/nosurf"
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

	engine.Use(NosurfHandler())
	engine.GET("/", index)
	engine.POST("/", postIndex)

	fmt.Println("Listening on http://127.0.0.1:8000/")
	engine.Run(":8000")
}

type nextHandler struct {
	next func(w http.ResponseWriter, r *http.Request)
}

func (h *nextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.next(w, r)
}

// Do not use this soultion
func NosurfHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := nosurf.New(nil)
		next := &nextHandler{
			next: func(w http.ResponseWriter, r *http.Request) {
				c.Request = r
				c.Next()
			},
		}
		h.SetSuccessHandler(next)

		h.ServeHTTP(c.Writer, c.Request)
	}
}
