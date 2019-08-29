package parallelhttp

// Structure to represent a Http Request
type HttpRequest struct {
	Url    string
	Method string
}

// Structure to represent a Http Response
type HttpResponse struct {
	Request    HttpRequest
	Body       string
	StatusCode int
	Error      error
}
