package errors

func NewParamsHydrationError(err error) *GogetaError {
	if err == nil {
		return nil
	}
	return &GogetaError{
		Status:  400,
		Message: "Invalid Parameters",
		Err:     err,
	}
}
