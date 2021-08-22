package service

import "github.com/stretchr/testify/mock"

type MockServiceError struct {
	mock.Mock
}

func (s *MockServiceError) Error() string {
	s.Called()
	return "mock"
}
