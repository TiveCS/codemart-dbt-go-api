package model

import "time"

type Response struct {
	Message      string `json:"message,omitempty"`
	Data         any    `json:"data,omitempty"`
	ProcessTime  string `json:"process_time,omitempty"`
	ProcessStart int64  `json:"process_start,omitempty"`
	ProcessEnd   int64  `json:"process_end,omitempty"`
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

func (r *Response) WithProcessTime(processStart int64, processEnd int64) *Response {
	r.ProcessTime = time.Duration(processEnd - processStart).String()
	r.ProcessStart = processStart
	r.ProcessEnd = processEnd
	return r
}
