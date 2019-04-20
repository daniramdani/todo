package helper

type (
	// TodoResponse
	TodoResponseStrc struct {
		RequestID    string      `json:"request_id"`
		Code         int         `json:"code"`
		ErrorMessage string      `json:"error_message,omitempty"`
		Data         interface{} `json:"data"`
	}
)

// TodoResponse - instantiate new Response
func TodoResponse() TodoResponseStrc {
	reqID := RandomStringBase64(20)
	return TodoResponseStrc{
		RequestID: reqID,
	}
}
