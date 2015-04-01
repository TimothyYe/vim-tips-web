package routers

import (
	"math/rand"
	"time"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
)

func getRandomIndex(total int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(total)
}

func HandleIndex(r render.Render) {
	tip := models.Tips{}

	db := GetDBInstance()
	total, err := db.C("tips").Count()

	if err != nil || total == 0 {
		r.HTML(200, "500", nil)
		return
	}

	index := getRandomIndex(total)

	if index == total {
		index = total - 1
	}

	db.C("tips").Find(nil).Skip(index).One(&tip)

	if db.Session != nil {
		defer db.Session.Close()
	}

	r.HTML(200, "index", map[string]interface{}{
		"Comment": tip.Comment,
		"Content": tip.Content,
		"Id":      tip.Id.Hex(),
		"IsIndex": true})
}
