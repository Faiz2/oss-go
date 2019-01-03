package company

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhCompany struct {
	Id          string        `json:"id"`
	Id_         bson.ObjectId `bson:"_id"`
	CompanyName string        `json:"companyname" bson:"company_name"`
}

func (bd *PhCompany) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhCompany) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhCompany) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhCompany) QueryId() string {
	return bd.Id
}

func (bd *PhCompany) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhCompany) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhCompany) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhCompany) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhCompany) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
