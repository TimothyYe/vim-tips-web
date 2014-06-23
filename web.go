package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

var (
	m = martini.Classic()
)

func initEnv() {

}

func Hello() string {
	return "Hello Server!"
}

func main() {
	initEnv()

	m.Get("/", Hello)
	http.Handle("/", m)

	http.ListenAndServe("3000", nil)
	fmt.Println("Server started...")
}
