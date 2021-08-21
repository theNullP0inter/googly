package resource

// DbResourceManager should be implemented by any resource manager that intends to manage a DB.
//
// DB can be either rdb or mongo_db. can be extended to others as well
type DbResourceManager interface {
	CrudResourceManager
}
