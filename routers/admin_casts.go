package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
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
