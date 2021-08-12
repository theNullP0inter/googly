package resource

type QueryParameters interface{}

type PaginationQueryParameters struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	OrderBy   string `form:"order_by"`
	OrderDesc bool   `form:"order_desc"`
}
