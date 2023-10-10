package exception

type HttpException struct {
	Status  int
	Message string
}

func ThrowHttpException(status int, message string) *HttpException {
	return &HttpException{status, message}
}
