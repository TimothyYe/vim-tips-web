package models

import (
	"labix.org/v2/mgo/bson"
)

type Identity struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Email    string
	Password []byte
}
