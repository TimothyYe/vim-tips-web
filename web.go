package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/timothyye/vim-tips-web/routers"
	"net/http"
)

var (
	m = martini.Classic()
)

func initEnv() {
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tpl", ".html"},
		Charset:    "UTF-8",
	}))

	routers.InitRouters()
}

func Hello(r render.Render) {
	r.HTML(200, "hello", "Timothy")
}

func main() {
	initEnv()

	m.Get("/", Hello)
	http.Handle("/", m)
	fmt.Println("Server started...")

	err := http.ListenAndServe("127.0.0.1:3001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
