/* package main

import (
	"fmt"
	"net/http"
)

type RequestURL struct {
	url        string
	statuscode int
}

func main() {

	c := make(chan RequestURL)
	results := make(map[string]int)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	fmt.Println("Checking URL ..........")
	for _, v := range urls {
		go hitURL(v, c)
	}

	for i := 0; i < len(urls); i++ {
		r := <-c
		results[r.url] = r.statuscode
	}

	for key, value := range results {
		fmt.Println("사이트 : ", key, "\n상태코드 : ", value)
	}

}

func hitURL(url string, c chan RequestURL) {

	resp, _ := http.Get(url)
	c <- RequestURL{url: url, statuscode: resp.StatusCode}
}
*/
