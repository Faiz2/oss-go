package auth

import (
	"oss-go/profile"

	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhAuth struct {
	Id      string            `json:"id"`
	Id_     bson.ObjectId     `bson:"_id"`
	Profile profile.PhProfile `json:"Profile" jsonapi:"relationships"`
	Token   string            `json:"token"`
}

func (bd *PhAuth) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhAuth) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhAuth) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhAuth) QueryId() string {
	return bd.Id
}

func (bd *PhAuth) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhAuth) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhAuth) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Profile":
		bd.Profile = v.(profile.PhProfile)
	}
	return bd
}

func (bd PhAuth) QueryConnect(tag string) interface{} {
	switch tag {
	case "profile":
		return bd.Profile
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhAuth) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhAuth) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhAuth) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
