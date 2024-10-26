package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	err, bodyData := RequestAndProvide()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Body: ", string(bodyData))
}

func RequestAndProvide() (error, []byte) {
	resp, err := http.Get("https://api.github.com/repos/aivruu/repo-viewer")
	if err != nil {
		return err, nil
	}
	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error during ReadCloser closing: ", err)
		}
	}(resp.Body)
	err = nil
	return err, readBytes
}
