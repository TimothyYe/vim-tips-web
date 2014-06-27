package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func Initialize(m *martini.ClassicMartini) {
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Directory:  "templates",
		Extensions: []string{".tpl", ".html"},
		Charset:    "UTF-8",
	}))

	InitRouters(m)
}

func InitRouters(m *martini.ClassicMartini) {
	m.Get("/", HandleIndex)
}
