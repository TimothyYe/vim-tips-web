package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func AdminShowTips(req *http.Request, r render.Render, db *mgo.Database, pager *paginate.Paginator) {
	num, _ := db.C("tips").Count()

	pers := 12
	pager.Init(pers, num)

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

func AdminAddTips(req *http.Request, r render.Render, db *mgo.Database) {
	content := req.FormValue("tip")
	comment := req.FormValue("comment")

	tip := models.Tips{Id: bson.NewObjectId(), Content: content, Comment: comment}
	db.C("tips").Insert(tip)

	r.Redirect("/admin/tips")
}

func AdminDelTips(req *http.Request, r render.Render, db *mgo.Database) {
	id := req.FormValue("Id")
	db.C("tips").RemoveId(bson.ObjectIdHex(id))

	r.Redirect("/admin/tips")
}

func AdminModifyTips(req *http.Request, r render.Render, db *mgo.Database) {
	r.HTML(200, "admin/index", map[string]interface{}{
		"IsTips": true}, render.HTMLOptions{Layout: "admin/layout"})
}
