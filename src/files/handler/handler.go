package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//上传页
		data, err := ioutil.ReadFile("./public/view/index.html") //先不使用模版方法
		if err != nil {
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收数据并且保存
		fmt.Println("upload")
	}
}
