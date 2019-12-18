package errhandle

import (
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, r *http.Request) error

//错误处理标准化函数
func ErrHandle(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)

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
