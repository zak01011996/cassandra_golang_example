package ctrls

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"testproj/acl"
	"testproj/models"

	"github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
	"github.com/zenazn/goji/web"
)

var (
	log *logrus.Logger // Place to store logger

	// Initializing errors
	authError       = errors.New("Login or password is incorrect")
	regError        = errors.New("Not enough data to reg user")
	postCreateError = errors.New("Not enough data to create post")
)

// This is response struct
type Resp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// This we need to prepare our msg to return
func prepareResp(isSuccess bool, msg string) string {
	msgStmt := Resp{isSuccess, msg}
	resp, err := json.Marshal(msgStmt)
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(resp)
}

// Create comment handler
func Index(c web.C, w http.ResponseWriter, r *http.Request) {
	// Initializing post
	post := models.Post{}

	// Getting all posts
	data := post.GetAll()

	// Prepairing success response
	resp, err := json.Marshal(data)
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	w.Write([]byte(resp))
}

// SignIn handler
func SignIn(c web.C, w http.ResponseWriter, r *http.Request) {

	// Creating new decoder from request body
	decoder := json.NewDecoder(r.Body)

	// Initializing dst user model
	user := models.User{}

	// Decoding to user model
	err := decoder.Decode(&user)
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	// Initializing sha1 hasher
	hasher := sha1.New()
	hasher.Write([]byte(user.Pass))
	user.Pass = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// Trying to authorize user
	err = user.Auth()
	if err != nil {
		ErrorRequest(authError, w)
		return
	}

	// Generating new token
	token := gocql.TimeUUID().String()

	// Initializing and saving session
	sess := models.Session{user.Id, user.Perm, user.Name, token}
	err = sess.Save()
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	// Prepairing success response
	resp := prepareResp(true, token)
	w.Write([]byte(resp))
}

// SignUp handler
func SignUp(c web.C, w http.ResponseWriter, r *http.Request) {

	// Creating new decoder from request body
	decoder := json.NewDecoder(r.Body)

	// Initializing dst user model
	user := models.User{}

	// Decoding to user model
	err := decoder.Decode(&user)
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	// Checking given data
	if user.Login == "" || user.Pass == "" || user.Name == "" {
		ErrorRequest(regError, w)
		return
	}

	// Initializing sha1 hasher
	hasher := sha1.New()
	hasher.Write([]byte(user.Pass))
	user.Pass = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// Setting user permissions
	user.Perm = acl.USER

	// Trying to reg user
	err = user.Reg()
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	// Prepairing success response
	resp := prepareResp(true, "User successfully registered")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resp))
}

// SignOut handler
func SignOut(c web.C, w http.ResponseWriter, r *http.Request) {

	// Trying to get session values
	token := r.Header.Get("auth_token")
	if token == "" {
		BadRequest(w)
		return
	}

	sess := models.Session{Token: token}
	err := sess.Delete()
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	resp := prepareResp(true, "")
	w.Write([]byte(resp))
}

// Get comment handler
func CreatePost(c web.C, w http.ResponseWriter, r *http.Request) {
	// Checking user access
	acs, sess, err := acl.CheckAccess(r, acl.USER)
	if err != nil || !acs {
		log.Error(err)
		BadRequest(w)
		return
	}

	// Creating new decoder from request body
	decoder := json.NewDecoder(r.Body)

	// Initializing dst user model
	post := models.Post{}

	// Decoding to post model
	err = decoder.Decode(&post)
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	// Checking given data
	if post.Title == "" || post.Msg == "" {
		ErrorRequest(postCreateError, w)
		return
	}

	// Setting author
	post.Author = sess.Uname
	post.AuthorId = sess.Uid

	// Trying to save post
	err = post.Save()
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	resp := prepareResp(true, "Post successfully created")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resp))
}

// Get comment handler
func DeletePost(c web.C, w http.ResponseWriter, r *http.Request) {
	// Checking user access
	acs, _, err := acl.CheckAccess(r, acl.ADMIN)
	if err != nil || !acs {
		log.Error(err)
		BadRequest(w)
		return
	}

	// Parsing params
	postId, err := gocql.ParseUUID(c.URLParams["id"])
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	// Deleting post
	post := models.Post{Id: postId}
	err = post.Delete()
	if err != nil {
		ErrorRequest(err, w)
		return
	}

	resp := prepareResp(true, "Post successfully deleted")
	w.Write([]byte(resp))
}

// This to handle errors
func ErrorRequest(err error, w http.ResponseWriter) {

	// Logging error
	log.Error(err)

	// Prepairing message
	resp := prepareResp(false, "500 Internal server error: "+err.Error())

	// Here we are returning error
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(resp))
}

// This is http not found request handler
func NotFoundRequest(c web.C, w http.ResponseWriter, r *http.Request) {

	// Prepairing message
	resp := prepareResp(false, "404 Not found")

	// Setting status code 404
	w.WriteHeader(http.StatusNotFound)

	// Here we are writing error
	w.Write([]byte(resp))
}

// This is http bad request handler
func BadRequest(w http.ResponseWriter) {
	// Setting status code 403
	w.WriteHeader(http.StatusForbidden)
	resp := prepareResp(false, "403 Forbidden")

	// Here we are returning error
	w.Write([]byte(resp))
}

func Init(lg *logrus.Logger) {
	// Setting logger
	log = lg
}
