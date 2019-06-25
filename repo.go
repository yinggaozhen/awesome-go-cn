package main

import (
	"os"
	"io/ioutil"
	"text/template"
	gfm "github.com/shurcooL/github_flavored_markdown"
)

type content struct {
	Body string
}

func main() {
	generateHTML()
}

func generateHTML() (err error) {
	// options
	readmePath := "./README.md"
	tplPath := "tmpl/tmpl.html"
	idxPath := "tmpl/index.html"
	input, _ := ioutil.ReadFile(readmePath)
	body := string(gfm.Markdown(input))
	c := &content{Body: body}
	t := template.Must(template.ParseFiles(tplPath))
	f, err := os.Create(idxPath)
	t.Execute(f, c)
	return
}