package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

//构建一个展示文件列表web服务，核心在标准化错误输出

func main() {
	http.HandleFunc("/list/",
		errWrapper(optFiles))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

//接着开始进行错误的标准处理【对错误信息转化映射成外部可见的err】
//首先将handler的标准方法再进行包装
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

//函数通过对刚刚封装的appHandler进行错误处理中间层
func errWrapper(
	handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		//进行未知可能出现panic的recover的处理
		//defer

		//实现了中间层的代理调用，并且实现异常捕获
		err := handler(writer, request)
		if err != nil {
			//自定义层面的错误[判定错误类型]
			if userErr, ok := err.(UserError); ok {
				//将错误的结果：例如用户填写信息不完整之类，外加标准状态码返回
				//如果状态码复杂的话，这里新增扩展判断即可
				http.Error(writer,
					userErr.Message(),
					http.StatusForbidden)

				return
			}

			//进行错误处理的标准化[映射处理：错误转化]
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound

				//接着可以完善
			default:
				code = http.StatusInternalServerError
			}
			//将错误结果输出
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

//新增自定义的错误类型
type UserError interface {
	error
	Message() string
}

//由于存在errWrapper函数，使得每一个handler本身有中间层封装，返回对应err结果
//在调用部分errWrapper(optFiles)，则符合了标准库的传入参数标准
//这里就实现了golang的中间层以及pipeline的设计模式
func optFiles(writer http.ResponseWriter,
	request *http.Request) error {
	thisPath := request.URL.Path[len("/list/"):] //查找对应请求的文件名
	path := GetCurrentPath() + "/retryGo/"
	fmt.Println(path) //
	file, err := os.Open(path + "/" + thisPath)
	if err != nil {
		//panic(err)
		//log.Warn()//进行日志标记
		return err
	}

	defer file.Close() //defer放在处理error之后的意义就是先识别是否打开了资源，如果没有打开的话，defer关闭也就没有意义

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}

//获得设置的path?
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//获得编译路径
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
