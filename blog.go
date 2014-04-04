package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Post struct {
	Title string
	Text  string
}

type Posts []Post

func main() {

	m := martini.Classic()

	fmt.Println("Environment: " + martini.Env)

	m.Use(render.Renderer())

	m.Get("/posts", PostsHandler)
	m.Run()
}

func PostsHandler(r render.Render) {

	posts := Posts{Post{"Olá", "Prazer em conhecê-los!"}}
	r.JSON(200, posts)
}
