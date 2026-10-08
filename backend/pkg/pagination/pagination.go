// pkg/pagination/pagination.go
package pagination

type Result[T any] struct {
	Data  []T   `json:"data"`
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Pages int64 `json:"pages"` // total / limit arredondado
}

func New[T any](data []T, total int64, page, limit int) Result[T] {
	pages := total / int64(limit)
	if total%int64(limit) > 0 {
		pages++
	}
	return Result[T]{
		Data:  data,
		Total: total,
		Page:  page,
		Limit: limit,
		Pages: pages,
	}
}
