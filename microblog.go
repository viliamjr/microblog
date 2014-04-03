package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func main() {

	session_store := sessions.NewCookieStore([]byte("A-Tonga-da-Mironga-do-Kabulete"))

	m := martini.Classic()

	fmt.Println("Environment: " + martini.Env)

	m.Use(sessions.Sessions("user_control", session_store))
	m.Use(render.Renderer())

	m.Get("/", Index)
	m.Post("/login", Login)
	m.Get("/logout", Logout)
	m.Run()
}

func Index(session sessions.Session, r render.Render) {

	var data = struct{ Username string }{}
	v := session.Get("username")
	if v != nil {
		data.Username = v.(string)
	}
	r.HTML(200, "index", data)
}

func Login(session sessions.Session, r *http.Request, w http.ResponseWriter) {

	session.Set("username", r.FormValue("username"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(session sessions.Session, r *http.Request, w http.ResponseWriter) {

	session.Delete("username")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
