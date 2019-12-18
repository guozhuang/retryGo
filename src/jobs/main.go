package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type picUrl map[string]bool

var ch = make(chan int)

//任务系统
func main() {
	fmt.Println("hello jobs!")

	//简单实现任务系统

	for i := 0; i < 101; i++ {
		go opt(i)
	}

	count := 101
	for range ch {
		count--
		if count == 0 {
			close(ch)
		}
	}
}

func opt(i int) {
	picUrlData := fetchFile("/Users/momo/Documents/imgs/pic_" + strconv.Itoa(i))

	//先做一个扫描文件，并且根据对应文件内容，下载图片到指定目录
	for k, _ := range picUrlData {
		downPic(k)
	}

	ch <- 1
}

func downPic(url string) {
	target := "/Users/momo/Documents/featur_img/"

	length := len("") + 6
	picName := url[length:]

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != 200 {
		log.Println("下载图片错误", url, err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	out, _ := os.Create(target + picName)

	io.Copy(out, bytes.NewReader(body))

	return
}

func fetchFile(fileName string) picUrl {
	result := make(map[string]bool)

	//扫描文件：

	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return result
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		result[string(a)] = true
	}

	return result
}
