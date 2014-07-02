package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func HandleTip(r render.Render, db *mgo.Database, params martini.Params) {
	tip := models.Tips{}

	db.C("tips").FindId(bson.ObjectIdHex(params["Id"])).One(&tip)

	r.HTML(200, "index", map[string]interface{}{
		"Comment": tip.Comment,
		"Content": tip.Content,
		"Id":      tip.Id.Hex()})
}
