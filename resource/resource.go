package resource

// Resource can be any entity you're trying to manage.
//
// Typically a database model.
//
// Sometimes, you have to maintain resources on a 3rd party service.
// Then, you can implement a "virtual" resource
type Resource interface {
}

// ResourceManager is a base interface for a resource manager
//
// ResourceManager acts like a proxy between the service and ORM( or mongo-driver).
// i.e, if you want to migrate your service from mongo to rdb, it'll expose similar APIs
type ResourceManager interface {
	GetResource() Resource
}

// DataInterface is interface for the data that's passed to & from resource managers
type DataInterface interface {
}
