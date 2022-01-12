package module

// Data 数据结构类型
type Data interface {
	// Valid 判断data是否有效
	Valid() bool
}

// Valid 判断request是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

// Valid 判断response是否有效
func (rsp *Response) Valid() bool {
	return rsp.httpRsp != nil && rsp.httpRsp.Body != nil
}

// Valid 判断item是否有效
func (item Item) Valid() bool {
	return item != nil
}