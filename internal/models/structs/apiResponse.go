package structs

type ApiResponseStruct struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ApiResponse(error bool, message string, data any) ApiResponseStruct {
	return ApiResponseStruct{
		Error:   error,
		Message: message,
		Data:    data,
	}
}
