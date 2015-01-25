package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func AdminShowCasts(r render.Render, db *mgo.Database, pager *paginate.Paginator) {
	num, _ := db.C("casts").Count()

	pers := 12
	pager.Init(pers, num)

	casts := []models.Casts{}
	viewCasts := []models.CastsView{}

	db.C("casts").Find(nil).Limit(pers).Skip(pager.Offset()).All(&casts)

	for _, t := range casts {
		viewCasts = append(viewCasts,
			models.CastsView{Id: t.Id.Hex(), Author: t.Author, AuthorUrl: t.AuthorUrl,
				VisitCount: t.VisitCount, Title: t.Title, Intro: t.Intro, ShowNotes: t.ShowNotes, Url: t.Url, LogoUrl: t.LogoUrl})
	}

	r.HTML(200, "admin/casts_index", map[string]interface{}{
		"IsCasts":   true,
		"Casts":     viewCasts,
		"Paginator": pager,
		"Num":       num}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminAddCastsPage(r render.Render, db *mgo.Database) {
	r.HTML(200, "admin/casts_add", map[string]interface{}{
		"IsCasts": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminAddCasts(req *http.Request, r render.Render, db *mgo.Database) {
	author := req.FormValue("author")
	aurhor_url := req.FormValue("authorurl")
	title := req.FormValue("title")
	intro := req.FormValue("intro")
	show_notes := req.FormValue("shownotes")
	url := req.FormValue("url")
	logo_url := req.FormValue("logourl")

	cast := models.Casts{Id: bson.NewObjectId(), Author: author, AuthorUrl: aurhor_url,
		VisitCount: 0, Title: title, Intro: intro, ShowNotes: show_notes, Url: url, LogoUrl: logo_url}
	db.C("casts").Insert(cast)

	r.Redirect("/admin/casts")
}

func AdminDelCasts(req *http.Request, r render.Render, db *mgo.Database) {
	id := req.FormValue("Id")
	db.C("casts").RemoveId(bson.ObjectIdHex(id))

	r.Redirect("/admin/casts")
}

func AdminModifyCasts(r render.Render, db *mgo.Database, params martini.Params) {

	cast := models.Casts{}

	db.C("casts").FindId(bson.ObjectIdHex(params["Id"])).One(&cast)

	r.HTML(200, "admin/casts_modify", map[string]interface{}{
		"IsCasts": true,
		"Id":      params["Id"],
		"Cast":    cast}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminUpdateCasts(r render.Render, db *mgo.Database, req *http.Request) {
	author := req.FormValue("author")
	author_url := req.FormValue("authorurl")
	title := req.FormValue("title")
	intro := req.FormValue("intro")
	show_notes := req.FormValue("shownotes")
	url := req.FormValue("url")
	logo_url := req.FormValue("logourl")

	db.C("casts").UpdateId(bson.ObjectIdHex(req.FormValue("id")),
		bson.M{"$set": bson.M{"author": author,
			"authorurl": author_url,
			"title":     title,
			"intro":     intro,
			"shownotes": show_notes,
			"url":       url,
			"logourl":   logo_url}})

	r.Redirect("/admin/casts")
}
