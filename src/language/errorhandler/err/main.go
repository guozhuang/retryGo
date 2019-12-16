package main

import (
	"net/http"
	"os"
)

//定义一个统一的处理handler，包含错误
type appHandler func(writer http.ResponseWriter, r *http.Request) error

//错误处理标准化函数
func errWarpper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request) //默认的golanghandler接口

		if err != nil {
			//对应于http的err进行标准化多态适配并且展示前端的标准化处理
			switch {
			case os.IsNotExist(err):
				http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}

			//标准化处理之后，将这里的错误情景覆盖
		}
	}
}

func main() {
	//做一个简单的http server
	http.HandleFunc("/testerr/",
		errWarpper(HanleTest))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic("server is err")
	}
}
