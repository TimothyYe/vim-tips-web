package routers

import (
	"github.com/codegangsta/martini-contrib/render"
)

func HandleAPI(r render.Render) {
	r.HTML(200, "api", map[string]interface{}{
		"IsAPI": true})
}
