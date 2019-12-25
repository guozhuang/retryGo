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
	"os/exec"
	"strconv"
	"strings"
)

type picUrl map[string]int

var ch = make(chan int)

const jobFileTarDir = "/Users/momo/Documents/imgs/"
const jobImgTarDir = "/Users/momo/Documents/feature_img/"

//任务系统
func main() {
	fmt.Println("hello jobs!")

	//简单实现任务系统
	//single()

	count := 50

	//生成对应协程的读取文件
	//cmdFiles(count)

	for i := 0; i < count; i++ {
		//先mk对应的目录
		mkDir(jobImgTarDir + strconv.Itoa(i) + "/")
		go opt(i)
	}

	for range ch {
		count--
		if count == 0 {
			close(ch)
		}
	}
}

func cmdFiles(count int) {
	catCmd := "cat " + jobFileTarDir + "pic.csv |wc -l"

	cmd := exec.Command("bash", "-c", catCmd)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("err:", err.Error())
	}

	lines := string(out)
	//fmt.Println(lines)
	//执行结果带一个换行
	lines = strings.Trim(lines, "\n")
	lines = strings.Trim(lines, " ") //前面还有空格
	lineCount, err := strconv.Atoi(lines)
	if err != nil {
		fmt.Println("str to int err", err.Error())
	}

	averageLines := lineCount / count
	index := 1
	end := 0

	for i := 0; i < count; i++ {
		index = end + 1
		indexLine := strconv.Itoa(index)
		end = index + averageLines
		endLine := strconv.Itoa(end)
		//sed shell
		sedCmd := "sed -n '" + indexLine + "," + endLine + "p' " + jobFileTarDir + "pic.csv >" + jobFileTarDir + strconv.Itoa(i) + ".csv"
		fmt.Println(sedCmd)
		//"sed -n '".$start_line.",".$end_line."p' ".$dir.$filename." > ".$dir.$newDir.'/'.$newDir.'_'.$i.'.csv'

		sedCmdR := exec.Command("bash", "-c", sedCmd)
		_, err := sedCmdR.CombinedOutput()

		if err != nil {
			fmt.Println("sed err:", err.Error())
		}
	}
}

func mkDir(dir string) {
	mkcmd := "mkdir " + dir

	mkCmd := exec.Command("bash", "-c", mkcmd)
	_, err := mkCmd.CombinedOutput()

	if err != nil {
		fmt.Println("sed err:", err.Error())
	}
}

func single() {
	//picUrlData := fetchFile("/Users/momo/Documents/feature/pic.csv")

	//先做一个扫描文件，并且根据对应文件内容，下载图片到指定目录
	/*for k, _ := range picUrlData {
		//fmt.Println(k)
		//downPic(k)
	}*/
}

func opt(i int) {
	picUrlData := fetchFile(jobFileTarDir + strconv.Itoa(i) + ".csv")

	//先做一个扫描文件，并且根据对应文件内容，下载图片到指定目录
	for k, _ := range picUrlData {
		downPic(k, i)
	}

	ch <- 1
}

func downPic(url string, jobId int) {
	target := jobImgTarDir + strconv.Itoa(jobId) + "/"

	picName := url[39:]

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
	result := make(map[string]int)

	//扫描文件：

	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return result
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	count := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		count++
		result[string(a)] = count
	}

	return result
}
