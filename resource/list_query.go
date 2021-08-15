package resource

// ListQueryParameters acts as an interface for any list queries to  a resource manager
//
// DB resource managers convert this ListQuery into db queries
type ListQuery interface{}

// PaginatedListQuery provides required parameters to implement a pagination for a ListQuery on a ResourceManager
type PaginatedListQuery struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	OrderBy   string `form:"order_by"`
	OrderDesc bool   `form:"order_desc"`
}
