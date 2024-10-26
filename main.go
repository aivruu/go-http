package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	bodyDataChannel := make(chan []byte)
	errChannel := make(chan error)
	go requestAndProvide(bodyDataChannel, errChannel)

	select {
	case bodyData := <-bodyDataChannel:
		// Body has been received.
		fmt.Println("Body: ", string(bodyData))
	case err := <-errChannel:
		fmt.Println("Error: ", err)
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}

func requestAndProvide(dataChannel chan []byte, errChannel chan error) {
	resp, err := http.Get("https://api.github.com/repos/aivruu/repo-viewer")
	if err != nil {
		errChannel <- err
		return
	}
	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errChannel <- err
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error during ReadCloser closing: ", err)
		}
	}(resp.Body)
	err = nil
	dataChannel <- readBytes
}
