package routers

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func DB() martini.Handler {
	fmt.Println("DB handler is called...1")

	return func() {
		fmt.Println("DB handler is called...2")
	}
}

func Initialize(m *martini.ClassicMartini) {
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Directory:  "templates",
		Extensions: []string{".tpl", ".html"},
		Charset:    "UTF-8",
	}))

	m.Use(DB)

	InitRouters(m)
}

func InitRouters(m *martini.ClassicMartini) {
	m.Get("/", HandleIndex)
}
