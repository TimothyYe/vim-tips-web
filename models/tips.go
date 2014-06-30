package models

import (
	"labix.org/v2/mgo/bson"
)

type Tips struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Content string
	Comment string
}
