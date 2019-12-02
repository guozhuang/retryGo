package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//上传页
		data, err := ioutil.ReadFile("./public/view/index.html") //先不使用模版方法
		if err != nil {
			fmt.Printf("fail get index.html:%s", err.Error())
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收数据并且保存
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("fail get err:%s", err.Error())
			return
		}

		defer file.Close()

		//上传文件的默认处理:进行迁移
		//还没有进行重命名
		newFile, err := os.Create("/tmp/" + head.Filename) //先存储到tmp目录下
		if err != nil {
			fmt.Printf("fail create file err:%s", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file) //将内存中的内容copy到对应的tmp目录下
		if err != nil {
			fmt.Printf("fail copy file err:%s", err.Error())
			return
		}

		//上传完成，将对应状态进行重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload finish")
}
