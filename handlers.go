package main

import (
	"html/template"
	"net/http"
	"path"
	"strings"
)

// NewLink ...
func NewLink(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "links", "new.html")
	// link := Link{}

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	links := allLinks()
	if err := tmpl.Execute(w, links); err != nil {
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

	link := Link{Link: form.Get("Link"), Description: form.Get("Description"), Tags: tags}
	link.insert()
	http.Redirect(w, r, "/new", 301)
}

// ListLinks ...
func ListLinks(w http.ResponseWriter, r *http.Request) {

}
