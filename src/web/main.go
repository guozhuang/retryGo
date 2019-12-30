package main

import (
	"net/http"
	"web/app/handle/errhandle"
	"web/app/handle/hello"
)

//写一个简单的http server的框架【带标准化错误输出、配置文件、以及连接处理、单元测试】
func main() {
	http.HandleFunc("/hello",
		errhandle.ErrHandle(hello.Hello))

	//

	err := http.ListenAndServe(":8881", nil)

	if err != nil {
		panic("server is err")
	}
}
