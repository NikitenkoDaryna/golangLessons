package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

//http://localhost:3000/page2
func main() {
	//url := "http://localhost:3000/page1"

	resp, err := http.Get("http://localhost:3000/page2?id=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	resp2, err := http.PostForm("http://localhost:3000/page1", url.Values{"name": {"Daryna"},"surname":{"Nikitenko"}, "age": {"19"},"sex":{"Female"},"email":{"darynadobro@gmail.com"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()
	io.Copy(os.Stdout, resp2.Body)

}
