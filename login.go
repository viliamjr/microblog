package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"io/ioutil"
	"microblog/core"
	"net/http"
)

type Post struct {
	Title string
	Text  string
}

type Posts []Post

func main() {

	session_store := sessions.NewCookieStore([]byte("A-Tonga-da-Mironga-do-Kabulete"))

	m := martini.Classic()

	fmt.Println("Environment: " + martini.Env)

	m.Use(sessions.Sessions("user_control", session_store))
	m.Use(render.Renderer())
	m.Use(core.Auth("user", "passwd"))

	m.Get("/", Index)
	m.Post("/login", Login)
	m.Get("/logout", Logout)
	m.Run()
}

func Index(session sessions.Session, r render.Render, auth core.AuthData) {

	var session_data = struct {
		Username  string
		PostsData Posts
	}{}

	v := session.Get("username")

	if v != nil {
		session_data.Username = v.(string)
	}

	req, _ := http.NewRequest("GET", "http://foojr.com/blog/posts", nil) //XXX handle error
	resp, _ := auth.CheckRequest(req)

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body) //XXX handle error

	var posts Posts
	json.Unmarshal(data, &posts) //XXX handle error
	session_data.PostsData = posts

	r.HTML(200, "index", session_data)
}

func Login(session sessions.Session, r *http.Request, render render.Render) {

	session.Set("username", r.FormValue("username"))
	render.Redirect("/")
}

func Logout(session sessions.Session, render render.Render) {

	session.Delete("username")
	render.Redirect("/")
}
