package acl

import (
	"net/http"
	"testproj/models"
)

// Permissions
// 0 - Guest (just view),
// 1 - Admin (all permissions),
// 2 - User  (view and create)
const (
	GUEST = iota
	ADMIN
	USER
)

// This func will check access to method from session
func CheckAccess(r *http.Request, perm int) (acs bool, sess models.Session, err error) {
	// Trying to get session values
	token := r.Header.Get("auth_token")
	if token == "" {
		return
	}

	// Initializing sessions
	sess = models.Session{Token: token}

	// Checking user permissions
	err = sess.Check()
	if err != nil {
		return
	}

	// Admin have all privileges
	if sess.Perm == ADMIN {
		acs = true
		return
	}

	// Checking for user permissions
	if sess.Perm == USER && perm == USER {
		acs = true
		return
	}

	return
}
