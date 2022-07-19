package main

import (
	"fmt"
	"html/template"

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

	engine.Use(adapter.Wrap(nosurf.NewPure))
	engine.GET("/", index)
	engine.POST("/", postIndex)

	fmt.Println("Listening on http://127.0.0.1:8000/")
	engine.Run(":8000")
}
