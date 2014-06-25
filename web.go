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
	routers.Init(m)
	http.Handle("/", m)

	fmt.Println("Server started...")

	err := http.ListenAndServe("127.0.0.1:3001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
