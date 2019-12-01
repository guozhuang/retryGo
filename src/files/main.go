package main

import (
	"files/handler"
	"fmt"
	"net/http"
)

func main() {
	//简单路由
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err := http.ListenAndServe("8080", nil)
	if err != nil {
		fmt.Errorf("cannot find server, err%s", err.Error())
	}
}
