package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func AdminShowCasts(req *http.Request, r render.Render, db *mgo.Database, pager *paginate.Paginator) {
	num, _ := db.C("casts").Count()

	pers := 12
	pager.Init(pers, num)

	tips := []models.Tips{}
	viewTips := []models.TipsView{}

	db.C("casts").Find(nil).Limit(pers).Skip(pager.Offset()).All(&tips)

	for _, t := range tips {
		viewTips = append(viewTips, models.TipsView{Id: t.Id.Hex(), Content: t.Content, Comment: t.Comment})
	}

	r.HTML(200, "admin/casts_index", map[string]interface{}{
		"IsCasts":   true,
		"Tips":      viewTips,
		"Paginator": pager,
		"Num":       num}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminAddCastsPage(req *http.Request, r render.Render, db *mgo.Database) {
	r.HTML(200, "admin/casts_add", map[string]interface{}{
		"IsCasts": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminAddCasts(req *http.Request, r render.Render, db *mgo.Database) {
	author := req.FormValue("author")
	title := req.FormValue("title")
	intro := req.FormValue("intro")
	show_notes := req.FormValue("shownotes")
	url := req.FormValue("url")
	logo_url := req.FormValue("logourl")

	cast := models.Casts{Id: bson.NewObjectId(), Author: author,
		VisitCount: 0, Title: title, Intro: intro, ShowNotes: show_notes, Url: url, LogoUrl: logo_url}
	db.C("casts").Insert(cast)

	r.Redirect("/admin/casts")
}
