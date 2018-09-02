package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
)

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

func allLinks() []Link {
	rows, err := db.Query("SELECT * FROM links")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	links := []Link{}
	for rows.Next() {
		link := new(Link)
		rows.Scan(&link.ID, &link.Link, &link.Description)

		links = append(links, *link)
	}

	return links
}

func main() {
	defer db.Close()

	http.HandleFunc("/new", NewLink)
	http.HandleFunc("/create", CreateLink)
	http.HandleFunc("/list", ListLinks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
