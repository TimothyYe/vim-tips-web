package routers

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/timothyye/martini-paginate"
	"labix.org/v2/mgo"
)

func GetDBInstance() *mgo.Database {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Failed to connect to mongo DB...")
	}

	return session.DB("vim_tips")
}

func InitDB() martini.Handler {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Failed to connect to mongo DB...")
		panic(err)
	}

	return func(res http.ResponseWriter, r *http.Request, c martini.Context) {
		c.Map(session.DB("vim_tips"))
	}
}

func Initialize(m *martini.ClassicMartini) {
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Directory:  "templates",
		Extensions: []string{".tpl", ".html"},
		Charset:    "UTF-8",
	}))

	store := sessions.NewCookieStore([]byte("vim-tips-secure"))
	m.Use(sessions.Sessions("my_session", store))

	InitConnections()
	m.Use(InitDB())
	InitRouters(m)
}

func InitRouters(m *martini.ClassicMartini) {
	//Routers for front pages
	m.Get("/", HandleIndex)
	m.Get("/tips/:Id", HandleTip)
	m.Get("/random_tips/txt", HandleRandomTxtTip)
	m.Get("/random_tips/json", HandleRandomJsonTip)
	m.Get("/casts", paginate.Handler, HandleCasts)
	m.Get("/casts/:Id", ShowCast)
	m.Get("/api", HandleAPI)
	m.Get("/tools", HandleTools)
	m.Get("/about", HandleAbout)
	m.Get("/admin/login", ShowLoginPage)
	m.Get("/admin/logout", HandleLogout)
	m.Post("/admin/login", HandleLogin)

	//Routers for admin panel
	m.Group("/admin", func(r martini.Router) {
		r.Get("/index", HandleAdminIndex)
		r.Get("/tips", paginate.Handler, AdminShowTips)
		r.Post("/tips/add", AdminAddTips)
		r.Post("/tips/del", AdminDelTips)
		r.Post("/tips/update", AdminModifyTips)
		r.Get("/casts", paginate.Handler, AdminShowCasts)
		r.Get("/password", AdminPassword)
		r.Post("/password", AdminUpdatePassword)
		r.Get("/casts/add", AdminAddCastsPage)
		r.Post("/casts", AdminAddCasts)
		r.Post("/casts/del", AdminDelCasts)
		r.Post("/casts/modify", AdminUpdateCasts)
		r.Get("/casts/modify/:Id", AdminModifyCasts)
	}, validateSession)

}
