package routers

import (
	"html/template"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo/bson"
)

func HandleCasts(r render.Render, pager *paginate.Paginator) {

	db := GetDBInstance()

	num, _ := db.C("casts").Count()

	pers := 6
	pager.Init(pers, num)

	casts := []models.Casts{}
	viewCasts := []models.CastsView{}

	db.C("casts").Find(nil).Limit(pers).Skip(pager.Offset()).All(&casts)

	for _, t := range casts {
		viewCasts = append(viewCasts,
			models.CastsView{Id: t.Id.Hex(),
				Author:     t.Author,
				AuthorUrl:  t.AuthorUrl,
				VisitCount: t.VisitCount,
				Title:      t.Title,
				LogoUrl:    t.LogoUrl,
				Intro:      template.HTML(t.Intro),
				ShowNotes:  template.HTML(t.ShowNotes),
				Url:        t.Url})
	}

	if db.Session != nil {
		defer db.Session.Close()
	}

	r.HTML(200, "casts", map[string]interface{}{
		"IsCasts":   true,
		"Casts":     viewCasts,
		"Paginator": pager,
		"Num":       num})
}

func ShowCast(r render.Render, params martini.Params) {
	db := GetDBInstance()
	cast := models.Casts{}

	db.C("casts").FindId(bson.ObjectIdHex(params["Id"])).One(&cast)

	viewCast := models.CastsView{Id: cast.Id.Hex(),
		Author:     cast.Author,
		AuthorUrl:  cast.AuthorUrl,
		VisitCount: cast.VisitCount,
		Title:      cast.Title,
		LogoUrl:    cast.LogoUrl,
		Intro:      template.HTML(cast.Intro),
		ShowNotes:  template.HTML(cast.ShowNotes),
		Url:        cast.Url}

	if db.Session != nil {
		defer db.Session.Close()
	}

	r.HTML(200, "show", map[string]interface{}{
		"IsCasts":  true,
		"ViewCast": viewCast})
}
