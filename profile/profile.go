package profile

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
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

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *Profile) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *Profile) QueryId() string {
	return bd.Id
}

func (bd *Profile) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *Profile) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *Profile) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *Profile) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *Profile) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
