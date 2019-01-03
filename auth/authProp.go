package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhAuthProp struct {
	Id        string        `json:"id"`
	Id_       bson.ObjectId `bson:"_id"`
	AuthID    string        `json:"auth_id" bson:"auth_id"`
	ProfileID string        `json:"profile_id" bson:"profile_id"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhAuthProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhAuthProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhAuthProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhAuthProp) QueryId() string {
	return bd.Id
}

func (bd *PhAuthProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhAuthProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhAuthProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhAuthProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhAuthProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhAuthProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhAuthProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
