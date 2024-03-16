package model

type HTTPResponse struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message,omitempty"`
	Total   int         `json:"total,omitempty"`
	Code    int         `json:"Code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"error,omitempty"`
}

func (h *HTTPResponse) SetSuccess(message string, code int, data interface{}) *HTTPResponse {
	h.Message = message
	h.Code = code
	h.Data = data
	return h
}
func (h *HTTPResponse) SetError(message string, code int, err error) *HTTPResponse {
	h.Message = message
	h.Code = code
	h.Err = err.Error()
	return h
}
