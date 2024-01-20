package types

type PaginationQuery struct {
	Page		int
	PageSize	int
	Offset		int
}

type PaginationResponse struct {
	TotalData		int64			`json:"total_data"`
	CurrentPage		int				`json:"current_page"`
	TotalPage		int				`json:"total_page"`
	PageSize		int				`json:"page_size"`
}