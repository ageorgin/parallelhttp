// Package parallelhttp contains utilities to make multiple Http call in parallel
package parallelhttp

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Returns as much response as request in argument doing call in parallel
func DoParallelHttpCall(requests map[string]HttpRequest) map[string]HttpResponse {
	numberResponsesReceived := 0
	// initialize response haspmap
	responses := make(map[string]HttpResponse, len(requests))
	// create Go channel to communicate asynchronously
	ch := make(chan HttpResponse)

	// for each request we start a go routine to make all calls asynchronously
	for _, request := range requests {
		go doHttpCall(request, ch)
	}

	// we wait until all responses are received
	for numberResponsesReceived < len(requests) {
		// waiting for receiving element from channel
		resp := <-ch
		responses[resp.Request.Url] = resp
		numberResponsesReceived += 1
	}

	return responses
}

func doHttpCall(request HttpRequest, ch chan HttpResponse) {
	res, err := http.Get(request.Url)
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	// sending HttpResponse to channel
	ch <- HttpResponse{request, string(responseBody)}
}
