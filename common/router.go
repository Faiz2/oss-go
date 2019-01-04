package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	aliossconfig "oss-go/config"

	bmconfig "github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gorilla/mux"
	uuid "github.com/hashicorp/go-uuid"
)

var (
	rt     *mux.Router
	o      sync.Once
	alioss aliossconfig.AliOssConfig
)

func Router() *mux.Router {
	o.Do(func() {
		rt = mux.NewRouter()

		rt.HandleFunc("/oss-upload", upload2OssFunc)
	})
	return rt
}

func upload2OssFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		errMsg := "upload request method error, please use POST."
		SimpleResponseForErr(errMsg, w)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		defer file.Close()
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			errMsg := "upload file key error, please use key 'file'."
			SimpleResponseForErr(errMsg, w)
			return
		}

		var bmRouter bmconfig.BMRouterConfig
		bmRouter.GenerateConfig()

		uuid, err := uuid.GenerateUUID()
		lsttmp := strings.Split(handler.Filename, ".")
		exname := lsttmp[len(lsttmp)-1]

		localDir := bmRouter.TmpDir + "/" + uuid + "." + exname // handler.Filename
		f, err := os.OpenFile(localDir, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("OpenFile error")
			fmt.Println(err)
			errMsg := "upload local file open error."
			SimpleResponseForErr(errMsg, w)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		result := map[string]string{
			"file": uuid,
		}
		AliOssUploadFunc(localDir, "test")
		// bmalioss.PushOneObject("bmsass", uuid, localDir)

		response := map[string]interface{}{
			"status": "ok",
			"result": result,
			"error":  "",
		}
		jso := jsonapiobj.JsResult{}
		jso.Obj = response
		enc := json.NewEncoder(w)
		enc.Encode(jso.Obj)
	}
}

func SimpleResponseForErr(errMsg string, w io.Writer) {
	response := map[string]interface{}{
		"status": 401.1,
		"result": errMsg,
		"error":  "client error",
	}
	jso := jsonapiobj.JsResult{}
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
}

func AliOssUploadFunc(path, target string) {
	alioss.GenerateConfig()

	client, err := oss.New(alioss.EndPoint, alioss.AccessKeyID, alioss.AccessKeySecret)

	if err != nil {
		fmt.Println("OSS Get Instance Error ====>", err)
	}

	bucket, err := client.Bucket(alioss.Name)

	if err != nil {
		fmt.Println("OSS Get Bucker Error ====>", err)
	}

	err = bucket.PutObjectFromFile(target, path)

	if err != nil {
		fmt.Println("OSS Push File Error ====>", err)
	}
}

// func AliOssDownloadFunc(localDir, )
