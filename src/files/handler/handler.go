package handler

import (
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//上传页
		ioutil.ReadFile("./public/view/index.html")
	} else if r.Method == "POST" {
		//接收数据并且保存
	}
}
