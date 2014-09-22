package routers

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
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
		"IsIndex": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminShowTips(req *http.Request, r render.Render, db *mgo.Database, s sessions.Session) {
	validateSession(r, s, "/admin/login")

	num, _ := db.C("tips").Count()

	pers := 8
	pager := NewPaginator(req, pers, num)

	tips := []models.Tips{}
	viewTips := []models.TipsView{}

	db.C("tips").Find(nil).Limit(pers).Skip(pager.Offset()).All(&tips)

	for _, t := range tips {
		viewTips = append(viewTips, models.TipsView{Id: t.Id.Hex(), Content: t.Content, Comment: t.Comment})
	}

	r.HTML(200, "admin/tips_index", map[string]interface{}{
		"IsTips":    true,
		"Tips":      viewTips,
		"Paginator": pager,
		"Num":       num}, render.HTMLOptions{Layout: "admin/layout"})
}
