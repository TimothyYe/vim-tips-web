package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
	//"labix.org/v2/mgo"
)

func HandleIndex(r render.Render) {
	tip := models.Tips{"", ":F", "Test comment."}
	//db.C("tips").Find(nil).One(&tip)

	r.HTML(200, "index", tip)
}
