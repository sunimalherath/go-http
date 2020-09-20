package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}

	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	// instead of creating a byte slice that to fit the elements of the Body, we can use the below;
	// io.Copy(os.Stdout, resp.Body)

	// using custom Writer;
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
