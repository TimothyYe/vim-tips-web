package routers

import (
	"labix.org/v2/mgo"
	"testing"
)

func TestHandleRandomTxtTip(t *testing.T) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		t.Error("Failed to connect to mongo DB...")
	} else {
		t.Log("Connected to mongo DB...")
	}

}
