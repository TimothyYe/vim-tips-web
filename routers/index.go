package routers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
	"math/rand"
	"time"
)

func getRandomIndex(total int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(total)
}

func HandleIndex(r render.Render, db *mgo.Database) {
	tip := models.Tips{}
	total, _ := db.C("tips").Count()
	index := getRandomIndex(total)
	db.C("tips").Find(nil).Skip(index).One(&tip)

	r.HTML(200, "index", tip)
}
