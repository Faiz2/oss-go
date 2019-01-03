package authothers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"oss-go/auth"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/bmrouter/bmoauth"
	"github.com/alfredyang1986/blackmirror/jsonapi"
)

type PhAuthGenerateToken struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhAuthGenerateToken) Exec() error {

	tmp := b.BrickInstance().Pr
	bmah := tmp.(auth.PhAuth)
	h := md5.New()
	io.WriteString(h, bmah.Id)

	token := fmt.Sprintf("%x", h.Sum(nil))

	bmah.Token = token
	b.BrickInstance().Pr = bmah

	err := bmoauth.PushToken(token)

	return err
}

func (b *PhAuthGenerateToken) Prepare(pr interface{}) error {
	req := pr.(auth.PhAuth)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhAuthGenerateToken) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhAuthGenerateToken) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhAuthGenerateToken) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PhAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhAuthGenerateToken) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PhAuth = b.BrickInstance().Pr.(auth.PhAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
