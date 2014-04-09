package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/viliamjr/microblog/core"
	"net/http"
)

type Post struct {
	Title string
	Text  string
}

type Posts []Post

var postsdb Posts

func main() {

	postsdb = Posts{Post{"Olá", "Prazer em conhecê-los!"}}

	session_store := sessions.NewCookieStore([]byte("A-Tonga-da-Mironga-do-Kabulete"))

	m := martini.Classic()

	fmt.Println("Environment: " + martini.Env)

	m.Use(sessions.Sessions("user_control", session_store))
	m.Use(render.Renderer())
	m.Use(core.Auth("user", "passwd"))

	m.Get("/posts", PostsHandler)
	m.Post("/post", PostHandler)
	m.Get("/register", RegisterHandler)
	m.Run()
}

func PostsHandler(r render.Render, auth core.AuthData) {

	if auth.Check() {
		r.JSON(200, postsdb)
	} else {
		r.JSON(http.StatusUnauthorized, struct{ Error string }{"Not Authorized"})
	}
}

func RegisterHandler(session sessions.Session, r render.Render) {

	v := session.Get("username")
	if v == nil {
		r.Redirect("/")
	} else {
		r.HTML(200, "register", nil)
	}
}

func PostHandler(session sessions.Session, r *http.Request, render render.Render) {

	v := session.Get("username")
	if v != nil {
		title := r.FormValue("title")
		text := r.FormValue("text")
		post := Post{title, text}
		postsdb = append(postsdb, post)
	}
	render.Redirect("/")
}
