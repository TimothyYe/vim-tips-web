package routers

import (
	"math/rand"
	"time"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
)

func getRandomIndex(total int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(total)
}

func HandleIndex(r render.Render, db *mgo.Database) {
	tip := models.Tips{}
	total, _ := db.C("tips").Count()
	index := getRandomIndex(total)

	if index == total {
		index = total - 1
	}

	db.C("tips").Find(nil).Skip(index).One(&tip)

	r.HTML(200, "index", map[string]interface{}{
		"Comment": tip.Comment,
		"Content": tip.Content,
		"Id":      tip.Id.Hex(),
		"IsIndex": true})
}
