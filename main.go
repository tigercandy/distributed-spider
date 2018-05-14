package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error code Status: ", resp.StatusCode)
		return
	}
	// 获取网页内容编码
	e := determineEncoding(resp.Body)
	// 中文转码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
