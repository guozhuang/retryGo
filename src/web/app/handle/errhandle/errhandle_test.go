package errhandle

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, r *http.Request) error {
	panic(500)
}

type testingErrUser string

func (e testingErrUser) Error() string {
	return e.Message() //err.Error内接口已经实现
}

func (e testingErrUser) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, r *http.Request) error {
	return testingErrUser(http.StatusText(http.StatusBadRequest))
}

func errIsNotExist(writer http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}

//测试errhandler
func TestErrHandle(t *testing.T) {
	tests := []struct {
		Handle  appHandler
		Code    int
		Message string
	}{
		{errPanic, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)},
		{errUserError, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)},
		{errIsNotExist, http.StatusNotFound, http.StatusText(http.StatusNotFound)},
	}

	for _, tt := range tests {
		re := ErrHandle(tt.Handle)
		//进行请求的构造
		resp := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.baidu.com", nil)

		re(resp, request)
		//进行验证
		verifyResponse(resp.Result(), tt.Code, tt.Message, t)
	}
}

//测试server:不用假参数测试，而是真实使用一个server进行测试
func TestErrHandleInServer(t *testing.T) {
	tests := []struct {
		Handle  appHandler
		Code    int
		Message string
	}{
		{errPanic, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)},
		{errUserError, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)},
		{errIsNotExist, http.StatusNotFound, http.StatusText(http.StatusNotFound)},
	}

	for _, tt := range tests {
		f := ErrHandle(tt.Handle)

		//不进行构造请求，而是直接运行server
		server := httptest.NewServer(http.HandlerFunc(f)) //将函数转化为interface

		resp, _ := http.Get(server.URL)
		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")

		if resp.StatusCode != tt.Code || body != tt.Message {
			t.Errorf("expect (%d, %s); "+"got (%d, %s)",
				tt.Code, tt.Message,
				resp.StatusCode, body)
		}
	}
}

//
func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); "+"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
