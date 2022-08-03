package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Listening on http://127.0.0.1:8100/")
	http.ListenAndServe(":8100", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]string)
	templ.Execute(w, context)
}

var templ = template.Must(template.New("t1").Parse(templateString))

var templateString string = `
<!doctype html>
<html>
<body>
<iframe src="http://srv.com/api" title="W3Schools Free Online Web Tutorials"></iframe>

</body>
</html>
`
