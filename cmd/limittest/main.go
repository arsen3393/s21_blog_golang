package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:8080/"
	for i := 0; i < 101; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Request status code: %d, %s\n", resp.StatusCode, resp.Status)
		resp.Body.Close()
		//time.Sleep(50 * time.Millisecond)
	}
}
