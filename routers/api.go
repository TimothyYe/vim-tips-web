package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
)

func HandleAPI(r render.Render) {
	r.HTML(200, "api", map[string]interface{}{
		"IsAPI": true})
}

func HandleRandomTxtTip(db *mgo.Database) string {
	tip := models.Tips{}
	total, _ := db.C("tips").Count()
	index := getRandomIndex(total)
	db.C("tips").Find(nil).Skip(index).One(&tip)

	return tip.Content + " " + tip.Comment
}
