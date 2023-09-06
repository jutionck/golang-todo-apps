package commons

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"net/http"
)

func SendCreatedResponse(c *gin.Context, data any, responseType string) {
	c.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:        http.StatusCreated,
			Description: responseType,
		},
		Data: data,
	})
}

func SendSingleResponse(c *gin.Context, data any, responseType string) {
	c.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:        http.StatusOK,
			Description: responseType,
		},
		Data: data,
	})
}

func SendPageResponse(c *gin.Context, data []any, responseType string, paging model.Paging) {
	c.JSON(http.StatusOK, &model.PagedResponse{
		Status: model.Status{
			Code:        http.StatusOK,
			Description: responseType,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendNoContent(c *gin.Context, responseType string) {
	c.String(http.StatusNoContent, responseType)
}

func SendErrorResponse(c *gin.Context, code int, errorMessage string) {
	c.AbortWithStatusJSON(code, &model.Status{
		Code:        code,
		Description: errorMessage,
	})
}

func SendFileResponse(c *gin.Context, fileName string, responseType string) {
	c.JSON(http.StatusOK, &model.FileResponse{
		Status: model.Status{
			Code:        http.StatusOK,
			Description: responseType,
		},
		FileName: fileName,
	})
}
