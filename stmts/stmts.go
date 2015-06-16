package stmts

/*
   This package we need to store CQL statements
*/

const (

	/*
	   ========================================
	   Base statements to initialize database
	   ========================================
	*/

	// This to create KEYSPACE for project
	CREATE_KSPACE = `
	    CREATE KEYSPACE IF NOT EXISTS test_proj
	    WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
	`

	// Users table statement
	CREATE_USERS = `
	    CREATE TABLE IF NOT EXISTS test_proj.users(
		    id UUID, 
		    name TEXT, 
		    login TEXT, 
		    password TEXT, 
		    perm INT,
		    PRIMARY KEY(id, login, password)
	    );
	`

	// Users table statement
	CREATE_SESSIONS = `
	    CREATE TABLE IF NOT EXISTS test_proj.sessions(
		    uid UUID, 
		    uname TEXT, 
		    perm INT, 
		    auth_token TEXT, 
		    PRIMARY KEY(auth_token)
	    );
	`

	// Comments create statement
	CREATE_POSTS = `
	    CREATE TABLE IF NOT EXISTS test_proj.posts (
		    id UUID, 
		    title TEXT, 
		    author TEXT, 
		    author_id UUID, 
		    msg TEXT, 
		    PRIMARY KEY(id)
	    );
	`

	// Create administrator
	CREATE_ADMIN = `
	    INSERT INTO test_proj.users (id, name, login, password, perm) 
	    VALUES (a55b4785-111b-11e5-90e0-30f9edaddb1c, 'Administrator', 'admin','0DPiKuNIrrVmD8IUCuw1hQxNqZc=', 1)
	`

	/*
	   ========================================
	   User statements
	   ========================================
	*/

	// Register new user
	USR_REG = `INSERT INTO test_proj.users (id, name, login, password, perm) VALUES (?, ?, ?, ?, ?)`

	// Auth user
	USR_AUTH = `SELECT id, name, perm FROM test_proj.users WHERE login = ? AND password = ? ALLOW FILTERING`

	/*
	   ========================================
	   Sessions statements
	   ========================================
	*/

	// Save session
	SESS_SAVE = `INSERT INTO test_proj.sessions (uid, perm, uname, auth_token) VALUES (?, ?, ?, ?)`

	// Check session
	SESS_CHCK = `SELECT uid, perm, uname FROM test_proj.sessions WHERE auth_token = ?`

	// Delete session
	SESS_DELETE = `DELETE FROM test_proj.sessions WHERE auth_token = ?`

	/*
	   ========================================
	   Posts statements
	   ========================================
	*/

	// Save post
	POST_SAVE = `INSERT INTO test_proj.posts (id, title, author, author_id, msg) VALUES (?, ?, ?, ?, ?)`

	// Get all posts
	POST_GETALL = `SELECT id, title, author, author_id, msg FROM test_proj.posts`

	// Delete post
	POST_DELETE = `DELETE FROM test_proj.posts WHERE id = ?`
)
