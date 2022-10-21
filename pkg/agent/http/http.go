package agent

import (
	"fmt"
	"net/http"
)

// HTTPRequest is a simple http implementation that can make a request to an endpoint with a given method, body, headers, and expected response code
func HTTPRequest(method string, endpoint string, body string, headers map[string]string, expectedResponseCode int) {
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != expectedResponseCode {
		panic(fmt.Sprintf("Expected %d, got %d", expectedResponseCode, resp.StatusCode))
	} else {
		fmt.Println("HTTP request successful")
	}
}
