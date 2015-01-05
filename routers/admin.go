package routers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/timothyye/martini-paginate"
	"github.com/timothyye/vim-tips-web/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func validateSession(r render.Render, s sessions.Session) {
	isLogin := s.Get("IsLogin")

	if isLogin == nil {
		fmt.Println("Not login...")
		r.Redirect("/admin/login")
	}
}

func AdminPassword(r render.Render, s sessions.Session) {
	r.HTML(200, "admin/password", map[string]interface{}{
		"IsPassword": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminUpdatePassword(req *http.Request, r render.Render, db *mgo.Database) {
	pass := req.FormValue("password")
	verify := req.FormValue("verify")

	if pass == verify {

		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		info, err := db.C("id").Upsert(bson.M{"email": "admin@vim-tips.com"}, bson.M{"$set": bson.M{"password": hashedPass, "email": "admin@vim-tips.com"}})

		if err == nil {
			if info.Updated >= 0 {
				r.HTML(200, "admin/password", map[string]interface{}{
					"IsPost":     true,
					"IsPassword": true,
					"IsSuccess":  true}, render.HTMLOptions{Layout: "admin/layout"})
			}
		}
	} else {
		r.HTML(200, "admin/password", map[string]interface{}{
			"IsPost":     true,
			"IsPassword": true,
			"IsSuccess":  false}, render.HTMLOptions{Layout: "admin/layout"})
	}
}

func ShowLoginPage(r render.Render, s sessions.Session) {

	isLogin := s.Get("IsLogin")

	if isLogin != nil {
		r.Redirect("/admin/index")
		return
	}

	r.HTML(200, "admin/login", map[string]interface{}{
		"IsAbout": true}, render.HTMLOptions{})
}

func HandleLogout(r render.Render, s sessions.Session) {
	s.Delete("IsLogin")
	r.Redirect("/admin/login")
}

func HandleLogin(req *http.Request, r render.Render, s sessions.Session, db *mgo.Database) {
	username := req.FormValue("username")
	pass := req.FormValue("password")

	id := models.Identity{}

	err := db.C("id").Find(bson.M{"email": "admin@vim-tips.com"}).One(&id)

	if err != nil {
		fmt.Println(err.Error())
	}

	if username == "admin@vim-tips.com" && bcrypt.CompareHashAndPassword(id.Password, []byte(pass)) == nil {
		fmt.Println("Login success!")
		s.Set("IsLogin", true)

		r.Redirect("/admin/index")
	} else {
		r.Redirect("/admin/login")
	}
}

func HandleAdminIndex(r render.Render, db *mgo.Database) {
	tips_count, err := db.C("tips").Count()

	if err != nil {
		tips_count = 0
	}

	r.HTML(200, "admin/index", map[string]interface{}{
		"IsIndex":   true,
		"TipsCount": tips_count}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminShowTips(req *http.Request, r render.Render, db *mgo.Database, pager *paginate.Paginator) {
	num, _ := db.C("tips").Count()

	pers := 12
	pager.Init(pers, num)

	tips := []models.Tips{}
	viewTips := []models.TipsView{}

	db.C("tips").Find(nil).Limit(pers).Skip(pager.Offset()).All(&tips)

	for _, t := range tips {
		viewTips = append(viewTips, models.TipsView{Id: t.Id.Hex(), Content: t.Content, Comment: t.Comment})
	}

	r.HTML(200, "admin/tips_index", map[string]interface{}{
		"IsTips":    true,
		"Tips":      viewTips,
		"Paginator": pager,
		"Num":       num}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminAddTips(req *http.Request, r render.Render, db *mgo.Database) {
	content := req.FormValue("tip")
	comment := req.FormValue("comment")

	tip := models.Tips{Id: bson.NewObjectId(), Content: content, Comment: comment}
	db.C("tips").Insert(tip)

	r.Redirect("/admin/tips")
}

func AdminDelTips(req *http.Request, r render.Render, db *mgo.Database) {
	id := req.FormValue("Id")
	db.C("tips").RemoveId(bson.ObjectIdHex(id))

	r.Redirect("/admin/tips")
}

func AdminModifyTips(req *http.Request, r render.Render, db *mgo.Database) {
	r.HTML(200, "admin/index", map[string]interface{}{
		"IsTips": true}, render.HTMLOptions{Layout: "admin/layout"})
}

func AdminShowCasts(req *http.Request, r render.Render, db *mgo.Database, pager *paginate.Paginator) {
	num, _ := db.C("casts").Count()

	pers := 12
	pager.Init(pers, num)

	tips := []models.Tips{}
	viewTips := []models.TipsView{}

	db.C("casts").Find(nil).Limit(pers).Skip(pager.Offset()).All(&tips)

	for _, t := range tips {
		viewTips = append(viewTips, models.TipsView{Id: t.Id.Hex(), Content: t.Content, Comment: t.Comment})
	}

	r.HTML(200, "admin/casts_index", map[string]interface{}{
		"IsCasts":   true,
		"Tips":      viewTips,
		"Paginator": pager,
		"Num":       num}, render.HTMLOptions{Layout: "admin/layout"})
}
