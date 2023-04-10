package model

type Response struct {
	Message      string         `json:"message,omitempty"`
	Data         map[string]any `json:"data,omitempty"`
	ResponseTime int64          `json:"response_time,omitempty"`
}

func NewResponse() *Response {
	return new(Response)
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) WithData(data map[string]any) *Response {
	r.Data = data
	return r
}

func (r *Response) WithResponseTime(responseTime int64) *Response {
	r.ResponseTime = responseTime
	return r
}
