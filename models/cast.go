package models

import (
	"labix.org/v2/mgo/bson"
)

type Casts struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Author     string
	VisitCount int
}
