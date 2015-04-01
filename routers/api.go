package routers

import (
	"encoding/json"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo/bson"
)

func HandleAPI(r render.Render) {
	db := GetDBInstance()
	txt_api := models.API{}
	json_api := models.API{}

	err := db.C("apis").Find(bson.M{"type": "txt"}).One(&txt_api)

	if err != nil {
		r.HTML(200, "500", nil)
		return
	}

	err = db.C("apis").Find(bson.M{"type": "json"}).One(&json_api)

	if err != nil {
		r.HTML(200, "500", nil)
		return
	}

	r.HTML(200, "api", map[string]interface{}{
		"IsAPI":       true,
		"TxtCounter":  txt_api.Count,
		"JsonCounter": json_api.Count})

	if db.Session != nil {
		defer db.Session.Close()
	}
}

func HandleRandomTxtTip() string {
	db := GetDBInstance()
	tip := models.Tips{}
	api := models.API{}
	total, _ := db.C("tips").Count()
	index := getRandomIndex(total)
	db.C("tips").Find(nil).Skip(index).One(&tip)
	err := db.C("apis").Find(bson.M{"type": "txt"}).One(&api)

	api.Count++

	if err != nil {
		db.C("apis").Insert(&models.API{Id: bson.NewObjectId(), Type: "txt", Count: 0})
	} else {
		db.C("apis").Update(bson.M{"type": "txt"}, bson.M{"type": "txt", "count": api.Count})
	}

	data, err := json.Marshal(api)

	if err == nil {
		sendAll(data)
	}

	if db.Session != nil {
		defer db.Session.Close()
	}

	return tip.Content + " " + tip.Comment
}

func HandleRandomJsonTip(r render.Render) {

	db := GetDBInstance()
	tip := models.Tips{}
	api := models.API{}
	total, _ := db.C("tips").Count()
	index := getRandomIndex(total)
	db.C("tips").Find(nil).Skip(index).One(&tip)
	err := db.C("apis").Find(bson.M{"type": "json"}).One(&api)

	api.Count++

	if err != nil {
		db.C("apis").Insert(&models.API{Id: bson.NewObjectId(), Type: "json", Count: 0})
	} else {
		db.C("apis").Update(bson.M{"type": "json"}, bson.M{"type": "json", "count": api.Count})
	}

	data, err := json.Marshal(api)

	if err == nil {
		sendAll(data)
	}

	if db.Session != nil {
		defer db.Session.Close()
	}

	r.JSON(200, tip)
}
