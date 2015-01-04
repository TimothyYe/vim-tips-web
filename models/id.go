package models

import (
	"labix.org/v2/mgo/bson"
)

type Id struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Email  string
	Passwd string
}
