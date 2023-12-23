package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn/"
	resq, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url error : %s", err)
		return
	}

	body, err := io.ReadAll(resq.Body)
	if err != nil {
		fmt.Printf("read body error : %s", err)
		return
	}

	fmt.Printf("%s", body)
}
