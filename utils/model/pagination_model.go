package model

type QueryParams struct {
	Query string
	Order string
	Sort  string
}

func (qp *QueryParams) IsSortValid() bool {
	return qp.Sort == "ASC" || qp.Sort == "DESC"
}

type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

type RequestQueryParams struct {
	QueryParams
	PaginationParam
}

type Paging struct {
	Page        int `json:"paging"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}
