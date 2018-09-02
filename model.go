package main

// Tag ...
type Tag struct {
	Name string
}

// Link ...
type Link struct {
	ID          int
	Link        string
	Description string
	Tags        []Tag
}
