package routers

import (
	"labix.org/v2/mgo"
	"testing"
	"errors"
)

func InitDB() (t *testing.T, db *mgo.Database, error) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		t.Error("Failed to connect to mongo DB...")
	} else {
		t.Log("Connected to mongo DB...")
	}

	return session.DB("vim_tips_test"), err
}

func TestHandleRandomTxtTip(t *testing.T) {
	db, err := InitDB()
}
