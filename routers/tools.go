package routers

import (
	"github.com/codegangsta/martini-contrib/render"
)

func HandleTools(r render.Render) {
	r.HTML(200, "tools", map[string]interface{}{
		"IsTools": true})
}
