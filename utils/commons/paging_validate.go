package commons

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"strconv"
)

func ValidateRequestQueryParams(c *gin.Context) (model.RequestQueryParams, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		return model.RequestQueryParams{}, fmt.Errorf("invalid page number")
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit <= 0 {
		return model.RequestQueryParams{}, fmt.Errorf("invalid limit value")
	}

	order := c.DefaultQuery("order", "id")
	sort := c.DefaultQuery("sort", "ASC")

	return model.RequestQueryParams{
		QueryParams: model.QueryParams{
			Order: order,
			Sort:  sort,
		},
		PaginationParam: model.PaginationParam{
			Page:  page,
			Limit: limit,
		},
	}, nil
}

func PagingValidate(requestQueryParams model.RequestQueryParams) (model.PaginationQuery, string) {
	var paginationQuery model.PaginationQuery
	paginationQuery = GetPaginationParams(requestQueryParams.PaginationParam)
	orderQuery := "id"
	if requestQueryParams.QueryParams.Order != "" && requestQueryParams.QueryParams.Sort != "" {
		sorting := "ASC"
		if requestQueryParams.QueryParams.Sort == "desc" {
			sorting = "DESC"
		}
		orderQuery = fmt.Sprintf("%s %s", requestQueryParams.QueryParams.Order, sorting)
	}
	return paginationQuery, orderQuery
}
