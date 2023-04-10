package model

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return new(Response)
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}
