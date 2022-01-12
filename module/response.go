package module

import "net/http"

// Response 响应包
type Response struct {
	// httpRsp HTTP响应
	httpRsp *http.Response
	// depth 请求深度
	depth uint32
}

// NewResponse 实例化Response
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{
		httpRsp: httpResp,
		depth: depth,
	}
}

// HTTPRsp 获取http response
func (rsp *Response) HTTPRsp() *http.Response {
	return rsp.httpRsp
}

// Depth 获取请求深度
func (rsp *Response) Depth() uint32 {
	return rsp.depth
}