package main

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorMessage(status int, message string) ErrorMsg {
	return ErrorMsg{
		Status:  status,
		Message: message,
	}
}

func (e ErrorMsg) Error() string {
	return e.Message
}
