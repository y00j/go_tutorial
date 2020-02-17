package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	w := customWriter{}
	io.Copy(w, resp.Body)

}

func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
