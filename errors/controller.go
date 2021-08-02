package errors

func NewParamsHydrationError(err error) *GooglyError {
	if err == nil {
		return nil
	}
	return &GooglyError{
		Status:  400,
		Message: "Invalid Parameters",
		Err:     err,
	}
}
