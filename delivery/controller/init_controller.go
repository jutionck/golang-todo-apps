package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/usecase"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"net/http"
)

type InitController struct {
	router  *gin.RouterGroup
	useCase usecase.UserUseCase
}

// InitData godoc
// @Summary      Init data
// @Description  get data
// @Tags         Init
// @Accept       json
// @Produce      json
// @Success      201  {object}  string
// @Router       /init [get]
func (i *InitController) initializeHandler(c *gin.Context) {
	if _, err := i.useCase.InitData(); err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	commons.SendSingleResponse(c, nil, "initialized")
}

func (i *InitController) Route() {
	i.router.GET(config.GetInit, i.initializeHandler)

}

func NewInitController(r *gin.RouterGroup, uc usecase.UserUseCase) *InitController {
	controller := InitController{
		router:  r,
		useCase: uc,
	}
	return &controller
}
