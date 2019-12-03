package handler

import (
	"encoding/json" //后面切换使用性能更好的json格式化库
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

//文件根据散列值进行下载
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form["filehash"][0]

	fileMeta := meta.GetFileMeta(fileHash)

	f, err := os.Open(fileMeta.Location)
	if err != nil {
		fmt.Printf("download open file err:%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("download read file err:%s", err.Error())
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Descrption", "attachment;filename=\""+fileMeta.FileName+"\"")
	w.Write(data)
}

//文件删除
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(fileHash)
	//todo:需要进行删除容错
	os.Remove(fileMeta.Location)

	meta.RemoveFileMeta(fileHash)
	w.WriteHeader(http.StatusOK)
}
