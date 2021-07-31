package service

type ServiceRequestInterface interface {
	Validate() *ServiceError
}
