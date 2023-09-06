package commons

import (
	"github.com/jutionck/golang-todo-apps/utils/model"
	"math"
	"strconv"
)

func GetPaginationParams(params model.PaginationParam) model.PaginationQuery {
	var page int
	var take int
	var skip int
	if params.Page > 0 {
		page = params.Page
	} else {
		page = 1
	}

	if params.Limit == 0 {
		cfg := New()
		n, _ := strconv.Atoi(cfg.Get("DEFAULT_ROWS_PER_PAGE"))
		take = n
	} else {
		take = params.Limit
	}

	skip = (page - 1) * take

	return model.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

func Paginate(page, limit, totalRows int) model.Paging {
	return model.Paging{
		Page:        page,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
		TotalRows:   totalRows,
		RowsPerPage: limit,
	}
}
