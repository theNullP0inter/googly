package errors

func NewResourceConversionError() *GooglyError {

	return &GooglyError{
		Status:  400,
		Message: "Invalid Resource",
		Err:     nil,
	}
}
