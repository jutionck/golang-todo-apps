package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/delivery/middleware"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/domain/dto"
	"github.com/jutionck/golang-todo-apps/usecase"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"net/http"
)

type UserController struct {
	uc             usecase.UserUseCase
	r              *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

// User godoc
// @Summary Create a user
// @Description Create a new user
// @Accept json
// @Produce json
// @Security 	 Bearer
// @Tags User
// @Param Body body dto.UserRequestDto true "New User"
// @Success 201 {object} domain.User
// @Router /users [post]
func (u *UserController) createHandler(c *gin.Context) {
	var payload dto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userData := &domain.User{
		Email:    payload.Email,
		Password: payload.Password,
	}
	if err := u.uc.RegisterNew(userData); err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	commons.SendCreatedResponse(c, payload, "OK")
}

// ListUser godoc
// @Summary      List users
// @Description  List users
// @Tags         User
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Success      200  {array}   domain.User
// @Router       /users [get]
func (u *UserController) listHandler(c *gin.Context) {
	paginationParam, err := commons.ValidateRequestQueryParams(c)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	users, paging, err := u.uc.FindAll(paginationParam)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range users {
		response = append(response, v)
	}
	commons.SendPageResponse(c, response, "OK", paging)
}

// GetUser godoc
// @Summary      Get user
// @Description  get user by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  domain.User
// @Router       /users/{id} [get]
func (u *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := u.uc.FindById(id)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commons.SendSingleResponse(c, user, "OK")
}

// User godoc
// @Summary Update a user
// @Description Update a new user
// @Accept json
// @Produce json
// @Security 	 Bearer
// @Tags User
// @Param Body body dto.UserRequestDto true "New User"
// @Success 201 {object} domain.User
// @Router /users [put]
func (u *UserController) updateHandler(c *gin.Context) {
	var payload dto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userData := &domain.User{
		Email:    payload.Email,
		Password: payload.Password,
	}
	if err := u.uc.UpdateData(userData); err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	commons.SendSingleResponse(c, payload, "OK")
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  delete user by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id   path      string  true  "User ID"
// @Success      204  {object}  string
// @Router       /users/{id} [delete]
func (u *UserController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := u.uc.DeleteData(id)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commons.SendNoContent(c, "OK")
}

func (u *UserController) Route() {
	u.r.POST("/users", u.authMiddleware.RequireToken("admin"), u.createHandler)
	u.r.GET("/users", u.authMiddleware.RequireToken("admin"), u.listHandler)
	u.r.GET("/users/:id", u.authMiddleware.RequireToken("admin"), u.getHandler)
	u.r.PUT("/users", u.authMiddleware.RequireToken("admin"), u.updateHandler)
	u.r.DELETE("/users/:id", u.authMiddleware.RequireToken("admin"), u.deleteHandler)
}

func NewUserController(uc usecase.UserUseCase, r *gin.RouterGroup, am middleware.AuthMiddleware) *UserController {
	controller := &UserController{
		uc:             uc,
		r:              r,
		authMiddleware: am,
	}
	return controller
}
