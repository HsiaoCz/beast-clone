package handlers

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (a AppError) Error() string {
	return a.Message
}

func NewAppError(code int, msg string) AppError {
	return AppError{
		Status:  code,
		Message: msg,
	}
}
