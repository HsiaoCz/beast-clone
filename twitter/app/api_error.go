package app

type ErrorResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorResp) Error() string {
	return e.Message
}

func NewErrorResp(status int, message string) ErrorResp {
	return ErrorResp{
		Status:  status,
		Message: message,
	}
}
