package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// Tag ...
type Tag struct {
	Name string
}

// Link ...
type Link struct {
	Name        string
	Description string
	Tags        []Tag
}

// NewLink ...
func NewLink(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "links", "new.html")
	link := Link{}

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, link); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateLink ...
func CreateLink(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form
	tags := []Tag{}

	rawTags := strings.Split(form.Get("Tags"), ",")

	for _, t := range rawTags {
		tags = append(tags, Tag{t})
	}
	link := Link{form.Get("Name"), form.Get("Description"), tags}
	fmt.Printf("%v", link)
}

// ListLinks ...
func ListLinks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/new", NewLink)
	http.HandleFunc("/create", CreateLink)
	http.HandleFunc("/list", ListLinks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
