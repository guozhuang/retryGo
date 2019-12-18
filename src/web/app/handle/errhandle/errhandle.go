package errhandle

import (
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, r *http.Request) error

type userError interface {
	error //这里使用了接口的组合形式
	Message() string
}

//错误处理标准化函数
func ErrHandle(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		//对panic的处理
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)

				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			//可展示前端的错误
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//对应于http的err进行标准化多态适配并且展示前端的标准化处理
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}
