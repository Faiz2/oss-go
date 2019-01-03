package profile

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhProfileProp struct {
	Id        string        `json:"id"`
	Id_       bson.ObjectId `bson:"_id"`
	ProfileID string        `json:"profileid" bson:"profile_id"`
	CompanyID string        `json:"companyid" bson:"company_id"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhProfileProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhProfileProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhProfileProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhProfileProp) QueryId() string {
	return bd.Id
}

func (bd *PhProfileProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhProfileProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhProfileProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhProfileProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhProfileProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhProfileProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhProfileProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
