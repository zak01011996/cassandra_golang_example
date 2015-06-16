package db

import (
	"testproj/stmts"

	"github.com/gocql/gocql"
)

// Here we'll store connection
var Sess *gocql.Session

// This is wrapper for gocql
func Connect(machines ...string) error {

	// Creating cluster for cassandra
	cluster := gocql.NewCluster(machines...)
	cluster.Consistency = gocql.Quorum

	// Initializing session
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	// Storring cassandra session
	Sess = session

	// Initializing database
	err = initDB()
	if err != nil {
		return err
	}

	return nil
}

// This we need to initialize database scheme
func initDB() error {

	// Creating KEYSPACE (if not exists)
	err := Sess.Query(stmts.CREATE_KSPACE).Exec()
	if err != nil {
		return err
	}

	// Creating Users (if not exists)
	err = Sess.Query(stmts.CREATE_USERS).Exec()
	if err != nil {
		return err
	}

	// Creating Sessions (if not exists)
	err = Sess.Query(stmts.CREATE_SESSIONS).Exec()
	if err != nil {
		return err
	}

	// Creating Posts (if not exists)
	err = Sess.Query(stmts.CREATE_POSTS).Exec()
	if err != nil {
		return err
	}

	// Adding administrator
	return Sess.Query(stmts.CREATE_ADMIN).Exec()
}
