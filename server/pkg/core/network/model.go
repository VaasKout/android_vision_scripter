package network

import "bytes"

// HTTPRequest ...
type HTTPRequest struct {
	URL      string
	Body     *bytes.Buffer
	Headers  map[string]string
	Method   string
	LogBody  bool
	LogReq   bool
	Insecure bool
}

// MultipartRequest ...
type MultipartRequest struct {
	URL       string
	ImagePath string
	Boundary  string
	Log       bool
}
