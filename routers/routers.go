package routers

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

func InitDB(m *martini.ClassicMartini) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Failed to connect to mongo DB...")
		panic(err)
	}

	m.Map(session.DB("vim_tips"))
}

func Initialize(m *martini.ClassicMartini) {
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Directory:  "templates",
		Extensions: []string{".tpl", ".html"},
		Charset:    "UTF-8",
	}))

	InitConnections()
	InitDB(m)
	InitRouters(m)
}

func InitRouters(m *martini.ClassicMartini) {
	m.Get("/", HandleIndex)
	m.Get("/tips/:Id", HandleTip)
	m.Get("/random_tips/txt", HandleRandomTxtTip)
	m.Get("/random_tips/json", HandleRandomJsonTip)
	m.Get("/api", HandleAPI)
	m.Get("/tools", HandleTools)
	m.Get("/about", HandleAbout)
	m.Get("/admin/login", HandleLogin)
}
