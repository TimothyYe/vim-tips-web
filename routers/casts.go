package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
)

func HandleCasts(r render.Render, db *mgo.Database, pager *paginate.Paginator) {
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
				Intro:      t.Intro,
				ShowNotes:  t.ShowNotes,
				Url:        t.Url})
	}

	r.HTML(200, "casts", map[string]interface{}{
		"IsCasts":   true,
		"Casts":     viewCasts,
		"Paginator": pager,
		"Num":       num})
}

func ShowCast(r render.Render) {
	r.HTML(200, "show", map[string]interface{}{
		"IsCasts": true,
		"Id":      1})
}
