package resource

// QueryParameters acts as an interface for any list queries to  a resource manager
type QueryParameters interface{}

// PaginationQueryParameters provides required parameters to implement a pagination query on a resource manager
type PaginationQueryParameters struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	OrderBy   string `form:"order_by"`
	OrderDesc bool   `form:"order_desc"`
}
