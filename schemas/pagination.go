package schemas

import "github.com/attapon-th/null"

// Pagination pagination response metadata
type Pagination struct {
	TotalRecords null.Int    `json:"total_records"`
	TotalPages   null.Int    `json:"total_pages"`
	CurrentPage  null.Int    `json:"current_page"`
	SizePage     null.Int    `json:"size_page"`
	PrevPage     null.String `json:"prev_page"`
	NextPage     null.String `json:"next_page"`
} //	@name	Pagination

// PaginationWithoutTotal pagination
// without total record and total pages for response metadata
type PaginationWithoutTotal struct {
	CurrentPage null.Int    `json:"current_page"`
	SizePage    null.Int    `json:"size_page"`
	PrevPage    null.String `json:"prev_page"`
	NextPage    null.String `json:"next_page"`
} //	@name	Pagination

// NewPagination create new pagination
// default size_page = 10
// default current_page = 1
func NewPagination(currentPage, sizePage int64) *Pagination {
	if currentPage < 1 {
		currentPage = 1
	}
	if sizePage < 1 {
		sizePage = 10
	}
	return &Pagination{
		CurrentPage:  null.NewInt(currentPage, true),
		SizePage:     null.NewInt(sizePage, true),
		TotalRecords: null.NewInt(0, true),
		TotalPages:   null.NewInt(0, true),
		NextPage:     null.NewString("", false),
		PrevPage:     null.NewString("", false),
	}
}

// NewPaginationFull create new pagination with total record
func NewPaginationFull(currentPage, sizePage, totalRecord int64) *Pagination {
	p := NewPagination(currentPage, sizePage)
	return p.SetTotalRecord(totalRecord)
}

// SetTotalRecord set total record and calculate total pages
// if total record < 1 return total_pages = null and total_record = null
func (p *Pagination) SetTotalRecord(totalRecord int64) *Pagination {
	sizePage := p.SizePage.ValueOrZero()
	p.TotalRecords = null.IntFrom(totalRecord)
	if sizePage < 1 {
		sizePage = 10
	}
	if totalRecord < 1 {
		p.TotalPages.SetValid(0)
		p.TotalRecords.SetValid(0)
		return p
	}

	totalPage := (totalRecord / sizePage)
	if totalRecord%sizePage > 0 {
		totalPage++
	}
	p.TotalPages = null.IntFrom(totalPage)

	return p
}

// GetLimitOffset get limit and offset
func (p *Pagination) GetLimitOffset() (limit int64, offset int64) {
	limit = p.SizePage.ValueOrZero()
	offset = (p.CurrentPage.ValueOrZero() - 1) * p.SizePage.ValueOrZero()
	return

}

// GetPaginationWithoutTotal get pagination without total field (total_record, total_pages)
func (p *Pagination) GetPaginationWithoutTotal() *PaginationWithoutTotal {
	return &PaginationWithoutTotal{
		CurrentPage: null.NewInt(p.CurrentPage.ValueOrZero(), p.CurrentPage.Valid),
		SizePage:    null.NewInt(p.SizePage.ValueOrZero(), p.SizePage.Valid),
		PrevPage:    null.NewString(p.PrevPage.ValueOrZero(), p.PrevPage.Valid),
		NextPage:    null.NewString(p.NextPage.ValueOrZero(), p.NextPage.Valid),
	}
}
