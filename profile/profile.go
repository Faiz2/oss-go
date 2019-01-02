package profile

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id       string        `json:"id"`
	Id_      bson.ObjectId `bson:"_id"`
	Username string        `json:"username" bson:"username"`
	// Password string        `json:"Company" jsonapi:"relationships"`
}

func (bd *Profile) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *Profile) RestId_WithID() {
	bmmodel.ResetId_WithID(bd)
}
