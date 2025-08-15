package view

type Viewer interface {
	GetID() uint
}

// 2) Optional common base struct
type View struct {
	ID uint
}

// Promote a method so any struct embedding View satisfies Viewer.
func (v View) GetID() uint { return v.ID }

type Paginated[T Viewer] struct {
	Items      []T
	PageNumber int64
	PageSize   int64
	TotalSize  int64
}

type PagedRequest struct {
	PageNumber int64
	PageSize   int64
}

func NewPagedRequest(pageNumber int64, pageSize int64) PagedRequest {
	return PagedRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
}
