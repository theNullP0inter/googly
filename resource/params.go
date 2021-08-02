package resource

type ParameterInterface interface{}

type PaginationParameters struct {
	Page      int    `json:"page,default=0"`
	PageSize  int    `json:"page_size,default=30"`
	OrderBy   string `json:"order_by,default=id"`
	OrderDesc bool   `json:"order_desc,default=false"`
}

type CrudListParameters struct {
	*PaginationParameters
}
