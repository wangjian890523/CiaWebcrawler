package module

import "net/http"

// Request 表示一个爬虫请求
type Request struct {
	//HTTP请求
	httpReq *http.Request

	// 请求的URL
	URL string

	// 请求的深度，从0开始
	depth uint32

	// 请求的方法 (GET, POST 等)
	Method string

	// 请求头
	Headers map[string][]string

	// 请求体
	Body []byte

	// 父请求的URL，可用于追踪请求链
	ParentURL string

	// 请求的优先级，数字越小优先级越高
	Priority int

	// 额外的请求参数
	Extras map[string]interface{}
}

// NewRequest 创建一个新的请求实例
func NewRequest(url string, depth uint32) *Request {
	return &Request{
		httpReq: nil,
		URL:     url,
		depth:   depth,
		Method:  "GET",
		Headers: make(map[string][]string),
		Extras:  make(map[string]interface{}),
	}
}

// SetMethod 设置请求方法
func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

// SetHeader 设置请求头
func (r *Request) SetHeader(key, value string) *Request {
	r.Headers[key] = []string{value}
	return r
}

// SetBody 设置请求体
func (r *Request) SetBody(body []byte) *Request {
	r.Body = body
	return r
}

// SetParentURL 设置父请求URL
func (r *Request) SetParentURL(parentURL string) *Request {
	r.ParentURL = parentURL
	return r
}

// SetPriority 设置请求优先级
func (r *Request) SetPriority(priority int) *Request {
	r.Priority = priority
	return r
}

// SetExtra 设置额外参数
func (r *Request) SetExtra(key string, value interface{}) *Request {
	r.Extras[key] = value
	return r
}

func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

//获取请求的深度
func (req *Request) Depth() uint32 {
	return req.depth
}
