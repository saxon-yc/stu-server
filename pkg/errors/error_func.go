package errorcode

type ErrorCodeString struct {
	code        int32
	method, msg string
}

func New(code int32, method, msg string) *ErrorCodeString {
	return &ErrorCodeString{code, method, msg}
}

func (e *ErrorCodeString) Error() string {
	return e.msg
}

func (e *ErrorCodeString) Code() int32 {
	return e.code
}

func (e *ErrorCodeString) Method() string {
	return e.method
}

func (e *ErrorCodeString) SetErr(msg string) {
	e.msg = msg
}
