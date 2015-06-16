package models

import (
	"testproj/db"
	"testproj/stmts"

	"github.com/gocql/gocql"
)

// Session model
type Session struct {
	Uid   gocql.UUID
	Perm  int
	Uname string
	Token string
}

// Session save method
func (m *Session) Save() error {
	// Saving session to database
	return db.Sess.Query(stmts.SESS_SAVE, m.Uid, m.Perm, m.Uname, m.Token).Exec()
}

// Session delete method
func (m *Session) Delete() error {
	// Deleting session from database
	return db.Sess.Query(stmts.SESS_DELETE, m.Token).Exec()
}

// This method we need to check session by auth_token
func (m *Session) Check() error {
	// Checking session by auth_token
	return db.Sess.Query(stmts.SESS_CHCK, m.Token).Consistency(gocql.One).Scan(
		&m.Uid, &m.Perm, &m.Uname,
	)
}
