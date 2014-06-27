package routers

import (
	"github.com/codegangsta/martini-contrib/render"
)

func HandleIndex(r render.Render) {
	r.HTML(200, "index", "Timothy")
}
