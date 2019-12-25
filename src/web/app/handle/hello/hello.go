package hello

import (
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) error {
	output := []byte("hello world!")
	w.Write(output)
	return nil
	/*_, err := os.Open("test123.txt")
	return err*/
}
