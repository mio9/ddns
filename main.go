package main

import (
	"fmt"
	"io"
	"net/http"
	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("Hello, World!")
	ip := getIp()
	fmt.Println(ip)
}

func getIp() string {
	resp, err := http.Get("https://myip.wtf/text")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
