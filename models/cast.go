package models

import (
	"html/template"
	"labix.org/v2/mgo/bson"
)

type Casts struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Author     string
	AuthorUrl  string
	VisitCount int
	Title      string
	LogoUrl    string
	Intro      string
	ShowNotes  string
	Url        string
}

type CastsView struct {
	Id         string
	Author     string
	AuthorUrl  string
	VisitCount int
	Title      string
	LogoUrl    string
	Intro      template.HTML
	ShowNotes  template.HTML
	Url        string
}
