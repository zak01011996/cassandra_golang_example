package models

import (
	"testproj/db"
	"testproj/stmts"

	"github.com/gocql/gocql"
)

// Post model
type Post struct {
	Id       gocql.UUID `json:"id"`
	Title    string     `json:"title"`
	Author   string     `json:"auth"`
	Msg      string     `json:"msg"`
	AuthorId gocql.UUID
}

// Get all posts
func (m *Post) GetAll() []Post {
	// Getting all posts
	result := make([]Post, 0)
	post := Post{}
	iter := db.Sess.Query(stmts.POST_GETALL).Iter()
	for iter.Scan(&post.Id, &post.Title, &post.Author, &post.AuthorId, &post.Msg) {
		result = append(result, post)
	}

	return result
}

// User registration method
func (m *Post) Save() error {
	// Saving user to database
	return db.Sess.Query(stmts.POST_SAVE, gocql.TimeUUID(), m.Title, m.Author, m.AuthorId, m.Msg).Exec()
}

// This method we need to delete post
func (m *Post) Delete() error {
	// Deleting post from database
	return db.Sess.Query(stmts.POST_DELETE, m.Id).Exec()
}
