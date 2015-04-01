package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo/bson"
)

func HandleTip(r render.Render, params martini.Params) {
	db := GetDBInstance()
	tip := models.Tips{}

	db.C("tips").FindId(bson.ObjectIdHex(params["Id"])).One(&tip)

	if db.Session != nil {
		defer db.Session.Close()
	}

	r.HTML(200, "tip", map[string]interface{}{
		"Comment": tip.Comment,
		"Content": tip.Content,
		"Id":      tip.Id.Hex()})
}
