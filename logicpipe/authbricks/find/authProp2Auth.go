package authfind

import (
	"fmt"
	"io"
	"net/http"
	"oss-go/auth"
	"oss-go/company"
	"oss-go/profile"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type PhAuthProp2AuthBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhAuthProp2AuthBrick) Exec() error {
	prop := b.bk.Pr.(auth.PhAuthProp)
	reval, err := findAuth(prop)
	profile, err := findProfile(prop)
	profileProp, err := findProfileProp(profile)
	company, err := findProfileCompany(profileProp)
	profile.Company = company
	reval.Profile = profile
	b.bk.Pr = reval
	return err
}

func (b *PhAuthProp2AuthBrick) Prepare(pr interface{}) error {
	req := pr.(auth.PhAuthProp)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhAuthProp2AuthBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhAuthProp2AuthBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhAuthProp2AuthBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PhAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhAuthProp2AuthBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PhAuth = b.BrickInstance().Pr.(auth.PhAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

/*------------------------------------------------
 * brick inner function
 *------------------------------------------------*/

func findAuth(prop auth.PhAuthProp) (auth.PhAuth, error) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.AuthID)
	req := request.Request{}
	req.Res = "PhAuth"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("Eqcond", condi)
	fmt.Println(c)

	reval := auth.PhAuth{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}

func findProfile(prop auth.PhAuthProp) (profile.PhProfile, error) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.ProfileID)
	req := request.Request{}
	req.Res = "PhProfile"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("Eqcond", condi)
	fmt.Println(c)

	reval := profile.PhProfile{}
	err := reval.FindOne(c.(request.Request))
	reval.Password = ""

	return reval, err
}

func findProfileProp(phProfile profile.PhProfile) (profile.PhProfileProp, error) {
	eq := request.Eqcond{}
	eq.Ky = "profile_id"
	eq.Vy = phProfile.Id
	req := request.Request{}
	req.Res = "PhProfileProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("Eqcond", condi)
	fmt.Println(c)

	reval := profile.PhProfileProp{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}

func findProfileCompany(prop profile.PhProfileProp) (company.PhCompany, error) {
	eq := request.Eqcond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.CompanyID)
	req := request.Request{}
	req.Res = "PhCompany"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("Eqcond", condi)
	fmt.Println(c)

	reval := company.PhCompany{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}
