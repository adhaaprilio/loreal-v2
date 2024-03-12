package errorHandler

type Error400 struct {
	Message string
}

type Error404 struct {
	Message string
}

type Error409 struct {
	Message string
}

type Error500 struct {
	Message string
}

func (e *Error400) Error() string {
	return e.Message
}
func (e *Error404) Error() string {
	return e.Message
}
func (e *Error409) Error() string {
	return e.Message
}
func (e *Error500) Error() string {
	return e.Message
}
