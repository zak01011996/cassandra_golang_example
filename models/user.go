package models

import (
	"testproj/db"
	"testproj/stmts"

	"github.com/gocql/gocql"
)

// User model
type User struct {
	Id    gocql.UUID `json:"id"`
	Name  string     `json:"name"`
	Login string     `json:"login"`
	Pass  string     `json:"pass"`
	Perm  int
}

// User registration method
func (m *User) Reg() error {
	// Saving user to database
	return db.Sess.Query(stmts.USR_REG, gocql.TimeUUID(), m.Name, m.Login, m.Pass, m.Perm).Exec()
}

// This method we need to make authorisation
func (m *User) Auth() error {
	// Checking user's login and password and getting his Id, Permissions and auth_token
	return db.Sess.Query(stmts.USR_AUTH, m.Login, m.Pass).Consistency(gocql.One).Scan(&m.Id, &m.Name, &m.Perm)
}
