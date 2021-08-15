package resource

// Resource can be any entity you're trying to manage.
//
// Typically a database model.
//
// Sometimes, you have to maintain resources on a 3rd party service.
// Then, you can implement a "virtual" resource
type Resource interface {
}

// DataInterface is interface for the data that's passed to & from resource managers
type DataInterface interface {
}
