package utils

type PaginationMeta struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_Size"`
	TotalPages int   `json:"total_pages"`
}
