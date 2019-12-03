package handler

import (
	"encoding/json"
	"files/meta"
	"files/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		//元数据处理:依旧不包含重命名机制
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location) //先存储到tmp目录下
		if err != nil {
			fmt.Printf("fail create file err:%s", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file) //将内存中的内容copy到对应的tmp目录下
		if err != nil {
			fmt.Printf("fail copy file err:%s", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		//上传完成，将对应状态进行重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload finish")
}

//根据filehash【通过sha1sum命令获取】获取fileMeta
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form["filehash"][0]

	fileMeta := meta.GetFileMeta(fileHash)

	data, err := json.Marshal(fileMeta)
	if err != nil {
		fmt.Printf("json encode err:%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
