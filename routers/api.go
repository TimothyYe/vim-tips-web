package routers

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"time"
)

func HandleAPI(r render.Render) {
	r.HTML(200, "api", map[string]interface{}{
		"IsAPI": true})

	go timer()
}

func timer() {
	for {
		fmt.Println("Time is:" + time.Now().String())
		time.Sleep(time.Second * 5)
		data := []byte(time.Now().String())

		fmt.Println("Send data from goroutine...")
		sendAll(data)
	}
}
