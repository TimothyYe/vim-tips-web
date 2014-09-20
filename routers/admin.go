package routers

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func validateSession(r render.Render, s sessions.Session, retUrl string) {
	isLogin := s.Get("IsLogin")

	if isLogin == nil {
		fmt.Println("Not login...")
		r.Redirect("/admin/login")
	} else {
		fmt.Println("Already login...")
	}
}

func ShowLoginPage(r render.Render) {
	r.HTML(200, "admin/login", map[string]interface{}{
		"IsAbout": true}, render.HTMLOptions{})
}

func HandleLogout(r render.Render, s sessions.Session) {
	s.Delete("IsLogin")
	r.Redirect("/admin/login")
}

func HandleLogin(req *http.Request, r render.Render, s sessions.Session) {
	username := req.FormValue("username")
	pass := req.FormValue("password")

	if username == "aa@bb.com" && pass == "111" {
		fmt.Println("Login success!")
		s.Set("IsLogin", true)

		r.Redirect("/admin")
	} else {
		r.Redirect("/admin/login")
	}
}

func HandleAdminIndex(r render.Render, s sessions.Session) {
	validateSession(r, s, "/admin/login")

	r.HTML(200, "admin/index", map[string]interface{}{
		"IsAbout": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminShowTips(r render.Render, s sessions.Session) {
	validateSession(r, s, "/admin/login")

	r.HTML(200, "admin/tips_index", map[string]interface{}{
		"IsAbout": true}, render.HTMLOptions{Layout: "admin/layout"})
}
