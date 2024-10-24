package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(GetRequest())
}

// should return error and string
func GetRequest() string {
	resp, err := http.Get("https://www.douyin.com/user/MS4wLjABAAAAKLbzdLrJxBLWIhuaDRJQYGV0sa7xmKvOovj_N0mRPhA")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		log.Fatal("mission failed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func PostRequest(url string) (string, error) {
	// request body
	// json
	body, err := json.Marshal(map[string]string{
		"key":   "hello",
		"value": "hi",
	})
	if err != nil {
		return "", err
	}
	// make a request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	// set request header
	// content type
	req.Header.Set("Content-Type", "application/json")

	// make a cliet to send request
	client := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err != nil {

		return "", err
	}
	// receive the response
	somebody, err := io.ReadAll(resp.Body)
	if err != nil {

		return "", err
	}
	// return the message
	return string(somebody), nil
}
