# Go ParallelHttp package
Package parallelhttp contains utilities to make multiple Http call in parallel

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installing
Installing with go get command line
```
go get -v github.com/ageorgin/parallelhttp
```

Cloning from github
```
cd $GOPATH
mkdir -p github.com/ageorgin/
git clone https://github.com/ageorgin/parallelhttp.git
```

### Using the package
```
package main

import (
	"fmt"

	"github.com/ageorgin/parallelhttp"
)

func main() {
	requests := map[string]parallelhttp.HttpRequest{
		"http://httpbin.org/get": parallelhttp.HttpRequest{"http://httpbin.org/get", "GET"},
		"http://www.google.fr":   parallelhttp.HttpRequest{"http://www.google.fr", "GET"},
	}

	responses := parallelhttp.DoParallelHttpCall(requests)
	fmt.Println(responses["http://httpbin.org/get"].Body)
	fmt.Println(responses["http://httpbin.org/get"].StatusCode)
}
```

## Authors

* **Arnaud Georgin** - *Initial work* - [ageorgin](https://github.com/ageorgin)