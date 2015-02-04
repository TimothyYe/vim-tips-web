package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/timothyye/vim-tips-web/routers"
	"net/http"
)

var (
	m = martini.Classic()
)

func main() {
	routers.Initialize(m)

	http.HandleFunc("/ws", routers.WSHandler)
	http.Handle("/", m)

	fmt.Println("Server started...")

	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowMessage(mess string) {
	fmt.Println(mess)
}
