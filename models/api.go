package models

import (
	"labix.org/v2/mgo/bson"
)

type API struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Type  string
	Count uint64
}
