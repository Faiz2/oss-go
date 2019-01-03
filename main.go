package main

import (
	"net/http"
	"oss-go/auth"
	"oss-go/company"
	authfind "oss-go/logicpipe/authbricks/find"
	others "oss-go/logicpipe/authbricks/others"
	"oss-go/profile"
	"sync"

	bmconfig "github.com/alfredyang1986/blackmirror/bmconfighandle"

	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmrouter"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
	fac.RegisterModel("Request", &request.Request{})
	fac.RegisterModel("Eqcond", &request.Eqcond{})
	fac.RegisterModel("Upcond", &request.Upcond{})
	fac.RegisterModel("Fmcond", &request.Fmcond{})
	fac.RegisterModel("BmErrorNode", &bmerror.BmErrorNode{})

	fac.RegisterModel("PhAuth", &auth.PhAuth{})
	fac.RegisterModel("PhAuthProp", &auth.PhAuthProp{})
	fac.RegisterModel("PhCompany", &company.PhCompany{})
	fac.RegisterModel("PhProfile", &profile.PhProfile{})
	fac.RegisterModel("PhProfileProp", &profile.PhProfileProp{})

	fac.RegisterModel("PhAuthFindProfileBrick", &authfind.PhAuthFindProfileBrick{})
	fac.RegisterModel("PhProfile2AuthProp", &authfind.PhProfile2AuthProp{})
	fac.RegisterModel("PhAuthProp2AuthBrick", &authfind.PhAuthProp2AuthBrick{})

	fac.RegisterModel("PhAuthGenerateToken", &others.PhAuthGenerateToken{})

	r := bmrouter.BindRouter()

	var once sync.Once
	var bmRouter bmconfig.BMRouterConfig
	once.Do(bmRouter.GenerateConfig)

	http.ListenAndServe(":"+bmRouter.Port, r)
}
