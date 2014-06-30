package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
)

func HandleIndex(r render.Render, db *mgo.Database) {
	tip := models.Tips{}
	db.C("tips").Find(nil).One(&tip)

	r.HTML(200, "index", tip)
}
