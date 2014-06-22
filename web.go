package main

import (
	//"fmt"
	"github.com/go-martini/martini"
	//"net/http"
)

var (
	m = martini.Classic()
)

func Hello() string {
	return "Hello Server!"
}

func main() {
	m.Get("/", Hello)
	m.Run()
}
