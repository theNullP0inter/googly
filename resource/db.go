package resource

type DbResourceManagerIntereface interface {
	CrudResourceManagerInterface
	GetModel() DataInterface
}
