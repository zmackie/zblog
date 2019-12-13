package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type Post struct {
	Body  string `json:"plaintext"`
	Title string `json:"title"`
	Date  string `json:"created_at"`
}
type Posts struct {
	Posts []Post `json:"posts"`
}

var jsonFile = "blog_posts.json"
var postsDir = "posts"

var postTemplate = `
+++
title = "{{.Title}}"
draft = false
date = "{{.Date}}"

+++
{{.Body}}
`

func mkFileName(title string) string {
	filename := fmt.Sprintf("%v.md", strings.Replace(title, " ", "-", -1))
	return path.Join("content", postsDir, filename)
}

func writeTemplate(tmpl *template.Template, p Post) error {
	filename := mkFileName(p.Title)

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, p)
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}

func main() {

	// read all the docs and put them each into a struct
	// - Pull the markdown field
	// - Pull the title field
	// - Pull the created at field
	// output all the structs to markdown files
	jf, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	// fmt.Println(jf)

	var posts Posts
	err = json.Unmarshal(jf, &posts)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s", posts)
	t := template.Must(template.New("post").Parse(postTemplate))

	for _, p := range posts.Posts {
		err = writeTemplate(t, p)
		if err != nil {
			fmt.Println("problem executing template:", err)
		}
	}

}
