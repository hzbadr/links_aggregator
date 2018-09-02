package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

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
	Link        string
	Description string
	Tags        []Tag
	ID          int
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:hassan@/links_aggregator")
	if err != nil {
		log.Fatal("hello", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal("world", err.Error())
	}
}

func (l *Link) insert() {
	result, err := db.Exec("INSERT INTO links (Link, Description) VALUES(?, ?)", l.Link, l.Description)

	if err != nil {
		log.Fatal("Cong", err.Error())
		return
	}

	i, _ := result.LastInsertId()

	l.ID = int(i)
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

	link := Link{Link: form.Get("Link"), Description: form.Get("Description"), Tags: tags}
	link.insert()
	http.Redirect(w, r, "/new", 301)
}

// ListLinks ...
func ListLinks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	defer db.Close()

	http.HandleFunc("/new", NewLink)
	http.HandleFunc("/create", CreateLink)
	http.HandleFunc("/list", ListLinks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
