package core

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"github.com/go-martini/martini"
	"net/http"
)

func Auth(username string, password string) martini.Handler {

	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		c.MapTo(&authData{username, password, res, req}, (*AuthData)(nil))
	}
}

type AuthData interface {
	Check() bool
	CheckRequest(*http.Request) (*http.Response, error)
}

type authData struct {
	user     string
	password string
	res      http.ResponseWriter
	req      *http.Request
}

func (t *authData) Check() bool {

	var siteAuth = base64.StdEncoding.EncodeToString([]byte(t.user + ":" + t.password))
	auth := t.req.Header.Get("Authorization")
	return secureCompare(auth, "Basic "+siteAuth)
}

func (t *authData) CheckRequest(req *http.Request) (*http.Response, error) {

	req.SetBasicAuth(t.user, t.password)
	client := &http.Client{}
	return client.Do(req) //XXX the buffer must be closed
}

func secureCompare(given string, actual string) bool {

	givenSha := sha256.Sum256([]byte(given))
	actualSha := sha256.Sum256([]byte(actual))
	return subtle.ConstantTimeCompare(givenSha[:], actualSha[:]) == 1
}
