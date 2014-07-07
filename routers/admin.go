package routers

import (
	"github.com/codegangsta/martini-contrib/render"
)

func HandleLogin(r render.Render) {
	r.HTML(200, "admin/login", map[string]interface{}{
		"IsAbout": true}, render.HTMLOptions{"admin/layout"})
}
