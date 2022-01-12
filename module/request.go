package module

import "net/http"

// Request 请求包
type Request struct {
	// httpReq HTTP请求
	httpReq *http.Request
	// depth 请求深度
	depth uint32
}

// NewRequest 实例化Request
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{
		httpReq: httpReq,
		depth: depth,
	}
}

// HTTPReq 获取http request
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// Depth 获取请求深度
func (req *Request) Depth() uint32 {
	return req.depth
}