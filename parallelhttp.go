// Package parallelhttp contains utilities to make multiple Http call in parallel
package parallelhttp

import (
	"io/ioutil"
	"net/http"
)

// Returns as much response as request in argument doing call in parallel
func DoParallelHttpCall(requests map[string]HttpRequest) map[string]HttpResponse {
	numberResponsesReceived := 0
	// initialize response haspmap
	responses := make(map[string]HttpResponse, len(requests))
	// create Go channel to communicate asynchronously
	ch := make(chan HttpResponse)
	// creating shared client between goroutines
	client := &http.Client{}
	// defer closing all connections after function returns
	defer client.CloseIdleConnections()

	// for each request we start a goroutine to make all calls asynchronously
	for _, request := range requests {
		go doHttpCall(client, request, ch)
	}

	// we wait until all responses are received
	for numberResponsesReceived < len(requests) {
		// waiting for receiving element from channel
		resp := <-ch
		responses[resp.Request.Url] = resp
		numberResponsesReceived += 1
	}

	// closing channel
	close(ch)

	return responses
}

func doHttpCall(client *http.Client, request HttpRequest, ch chan HttpResponse) {
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		ch <- HttpResponse{request, "", 400, err}
	}

	res, err := client.Do(req)
	if err != nil {
		ch <- HttpResponse{request, "", res.StatusCode, err}
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ch <- HttpResponse{request, "", 422, err}
	}
	res.Body.Close()

	// sending HttpResponse to channel
	ch <- HttpResponse{request, string(responseBody), res.StatusCode, nil}
}
