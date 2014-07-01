package routers

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
)

func HandleTip(r render.Render, db *mgo.Database, params martini.Params) {
	tip := models.Tips{}
	total, _ := db.C("tips").Count()
	index := getRandomIndex(total)
	db.C("tips").Find(nil).Skip(index).One(&tip)

	id := params["Id"]
	fmt.Println("Id is:" + id)

	r.HTML(200, "index", tip)
}
