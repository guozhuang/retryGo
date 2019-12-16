package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	Ua      string
	TimeOUt time.Duration
}

//这样就通过实现对应的逻辑，调用者通过声明接口，
// 从而将一些逻辑invoke到相应的被调用者中，
//进而实现了比较复杂的结构的interface的组织形式
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
