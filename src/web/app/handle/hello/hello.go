package hello

import (
	"net/http"
	"os"
)

func Hello(w http.ResponseWriter, r *http.Request) error {
	/*output := []byte("hello world!")
	w.Write(output)*/

	_, err := os.Open("test123.txt")
	return err
}
