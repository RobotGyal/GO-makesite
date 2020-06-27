package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

func main() {
	readFile()
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(fileContents))
	return string(fileContents)
}

func renderTemplate() {
	paths := []string{
		"template.tmpl",
	}

	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
	err = t.Execute(os.Stdout, makesite)
	if err != nil {
		panic(err)
	}
}
