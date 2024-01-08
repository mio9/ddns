package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	// "gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("Hello, World!")
	pwd, err := os.Getwd();
	data, err := os.ReadFile(path.Join(pwd, "config.yaml"))
	if err != nil {
		fmt.Println("Error reading config file")
		panic(err)
	}
	fmt.Println(string(data))


	// ip := getIp()
	// fmt.Println(ip)
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
