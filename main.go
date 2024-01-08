package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	ip := getIp()

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
