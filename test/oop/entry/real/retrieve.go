package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retrieve struct {
	//实际rpc调用时的处理
	UserAgent string
	TimeOut   time.Duration
}

//实际中的方法
func (r *Retrieve) Get(url string) string {
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
