package main

import (
	"files/handler"
	"fmt"
	"net/http"
)

func main() {
	//简单路由
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/getFileMeta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Errorf("cannot find server, err%s", err.Error())
	}
}
