package main

import (
	"net/http"
	"os"
)

//appHandler接口
func HanleTest(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/testerr/"):]
	if path == "123" {

		//手动设置一个错误状态,这样手动设置标准化
		/*http.Error(
		w,
		"page is not",
		http.StatusInternalServerError)*/

		//通过这里的err到errwarpper内进行接收
		_, err := os.Open("test123.txt")
		//fmt.Println(err)//验证没问题
		return err //直接将error return，由统一的错误处理来统一处理
	}

	return nil
}
