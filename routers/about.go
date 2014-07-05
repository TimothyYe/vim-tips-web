package routers

import (
	"github.com/codegangsta/martini-contrib/render"
)

func HandleAbout(r render.Render) {
	r.HTML(200, "about", map[string]interface{}{
		"IsAbout": true})
}
