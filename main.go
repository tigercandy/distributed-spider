package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error code Status:", resp.StatusCode)
		return
	}
	all, ok := ioutil.ReadAll(resp.Body)
}
