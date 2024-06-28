package handlers

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorMsg) Error() string {
	return e.Message
}

func ErrorMessage(code int, msg string) ErrorMsg {
	return ErrorMsg{
		Status:  code,
		Message: msg,
	}
}
