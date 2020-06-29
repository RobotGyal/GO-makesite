package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
)

type post struct {
	Author  string
	Content string
}

func readFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContents))
	return string(fileContents)
}

func renderTemplate(content string) string {
	paths := []string{
		"template.tmpl",
	}

	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, post{Author: "Aleia", Content: "Random Content for latest post"})
	if err != nil {
		panic(err)
	}
	fmt.Println(buffer.String())
	return buffer.String()
}

func writeFile(buffer string, file string) bool {
	bytes := []byte(buffer)
	err := ioutil.WriteFile(file, bytes, 0644)
	if err != nil {
		return false
	}
	return true
}

func main() {
	// FlAGS
	filePtr := flag.String("file", "latest-post.txt", "displays file content")
	flag.Parse()

	// readFile("first-post.txt")
	template := renderTemplate(*filePtr)
	namedFile := strings.SplitN(*filePtr, ".", 2)[0] + ".html"
	writeFile(template, namedFile)
}
