package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	tp, err := loadTemplates("fluxble", "./web")
	if err != nil {
		panic(err)
	}
	t = tp

	mux := http.NewServeMux()
	mux.HandleFunc("/auth", authFstHandler)
	mux.HandleFunc("/auth/local", authHandler)
	fmt.Println("Now server is running on port 3000.")
	http.ListenAndServe(":3000", mux)
}

func authFstHandler(w http.ResponseWriter, req *http.Request) {
	url := "/auth/local?req=abc"
	http.Redirect(w, req, url, http.StatusFound)
}

func authHandler(w http.ResponseWriter, req *http.Request) {

	data := struct {
		URL    string
		Method string
	}{"http://www.baidu.com", "get"}
	values := renderValues{
		"k1": "v1",
		"k2": "v2",
		"d":  data,
	}
	if err := t.auth(w, values); err != nil {
		w.Write([]byte(err.Error()))
	}
}

var (
	t *templates
)

type templates struct {
	theme     string
	templates map[string]*template.Template
}

func loadTemplates(theme, templatesDir string) (*templates, error) {
	templatesDir = filepath.Join(templatesDir, theme)
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %v", err)
	}

	filenames := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filenames = append(filenames, filepath.Join(templatesDir, file.Name()))
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("no files in template dir %q", templatesDir)
	}

	funcs := map[string]interface{}{
		"theme": func() string { return theme },
		"lower": strings.ToLower,
	}

	tmpls, err := template.New("").Funcs(funcs).ParseFiles(filenames...)
	if err != nil {
		return nil, fmt.Errorf("parse files: %v", err)
	}

	t := &templates{
		theme:     theme,
		templates: make(map[string]*template.Template, len(filenames)),
	}

	for _, fullname := range filenames {
		fn := filepath.Base(fullname)
		t.templates[fn] = tmpls.Lookup(fn)
	}

	return t, nil
}

type renderValues map[string]interface{}

func (t *templates) auth(w http.ResponseWriter, values renderValues) error {
	tmpl, ok := t.templates["index.html"]
	if !ok || tmpl == nil {
		return fmt.Errorf("missing template(s): %s", "index.html")
	}

	return t.renderTemplate(w, tmpl, values)
}

func (t *templates) renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) error {
	wr := w
	if err := tmpl.Execute(wr, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	return nil
}
