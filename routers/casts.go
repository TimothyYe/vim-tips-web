package routers

import (
	"github.com/codegangsta/martini-contrib/render"
)

func HandleCasts(r render.Render) {
	r.HTML(200, "casts", map[string]interface{}{
		"IsCasts": true})
}

func ShowCast(r render.Render) {
	r.HTML(200, "show", map[string]interface{}{
		"IsCasts": true,
		"Id":      1})
}
