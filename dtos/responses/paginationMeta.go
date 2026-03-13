package responses

type PaginationMeta struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_Size"`
	TotalPages int   `json:"total_pages"`
}

const (
	DefaultPage     = 1
	DefaultPageSize = 20
)

func NormalisePage(page, pageSize int) (int, int) {
	if page < 1 {
		page = DefaultPage
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = DefaultPageSize
	}
	return page, pageSize
}

func Offset(page, pageSize int) int {
	if page < 1 {
		page = DefaultPage
	}
	return (page - 1) * pageSize
}

func CalculateTotalPages(total int64, pageSize int) int {
	if pageSize == 0 {
		return 0
	}
	return int((total + int64(pageSize) - 1) / int64(pageSize))
}

func BuildMeta(total int64, page, pageSize int) PaginationMeta {
	page, pageSize = NormalisePage(page, pageSize)
	return PaginationMeta{
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: CalculateTotalPages(total, pageSize),
	}
}
