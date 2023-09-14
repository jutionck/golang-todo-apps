package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/delivery/middleware"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/domain/dto"
	"github.com/jutionck/golang-todo-apps/usecase"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"net/http"
)

type TodoController struct {
	uc             usecase.TodoUseCase
	r              *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

// Todo godoc
// @Summary Create a todo
// @Description Create a new todo
// @Accept json
// @Produce json
// @Security 	 Bearer
// @Tags Todo
// @Param Body body dto.TodoRequestDto true "New Todo"
// @Success 201 {object} domain.Todo
// @Router /todos [post]
func (u *TodoController) createHandler(c *gin.Context) {
	var payload dto.TodoRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	values, _ := c.Get("todo")
	todo := values.(jwt.MapClaims)
	todoData := &domain.Todo{
		IsCompleted: payload.IsCompleted,
		Name:        payload.Name,
		UserID:      todo["id"].(string),
	}
	if err := u.uc.RegisterNew(todoData); err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	commons.SendCreatedResponse(c, payload, "OK")
}

// ListTodo godoc
// @Summary      Todo users
// @Description  Todo users
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Success      200  {array}   domain.Todo
// @Router       /todos [get]
func (u *TodoController) listHandler(c *gin.Context) {
	paginationParam, err := commons.ValidateRequestQueryParams(c)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	values, exists := c.Get("todo")
	todo := values.(jwt.MapClaims)
	if !exists {
		commons.SendErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	var excludeID string
	if todo["role"] == "user" {
		excludeID = todo["id"].(string)
	}

	users, paging, err := u.uc.FindAll(paginationParam, excludeID)
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

// GetTodo godoc
// @Summary      Get todo
// @Description  get todo by ID
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  domain.Todo
// @Router       /todos/{id} [get]
func (u *TodoController) getHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := u.uc.FindById(id)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commons.SendSingleResponse(c, user, "OK")
}

// Todo godoc
// @Summary Update a todo
// @Description Update a new todo
// @Accept json
// @Produce json
// @Security 	 Bearer
// @Tags Todo
// @Param Body body dto.TodoRequestDto true "New Todo"
// @Success 201 {object} domain.Todo
// @Router /todos [put]
func (u *TodoController) updateHandler(c *gin.Context) {
	var payload dto.TodoRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todoData := &domain.Todo{
		IsCompleted: payload.IsCompleted,
		Name:        payload.Name,
	}
	if err := u.uc.UpdateData(todoData); err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	commons.SendSingleResponse(c, payload, "OK")
}

// DeleteTodo godoc
// @Summary      Delete todo
// @Description  delete todo by ID
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id   path      string  true  "Todo ID"
// @Success      204  {object}  string
// @Router       /todos/{id} [delete]
func (u *TodoController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := u.uc.DeleteData(id)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commons.SendNoContent(c, "OK")
}

func (u *TodoController) Route() {
	u.r.POST(config.PostTodo, u.authMiddleware.RequireToken(config.ADMIN, config.USER), u.createHandler)
	u.r.GET(config.GetTodoList, u.authMiddleware.RequireToken(config.ADMIN, config.USER), u.listHandler)
	u.r.GET(config.GetTodo, u.authMiddleware.RequireToken(config.ADMIN, config.USER), u.getHandler)
	u.r.PUT(config.PutTodo, u.authMiddleware.RequireToken(config.ADMIN, config.USER), u.updateHandler)
	u.r.DELETE(config.DeleteTodo, u.authMiddleware.RequireToken(config.ADMIN, config.USER), u.deleteHandler)
}

func NewTodoController(uc usecase.TodoUseCase, r *gin.RouterGroup, am middleware.AuthMiddleware) *TodoController {
	controller := &TodoController{
		uc:             uc,
		r:              r,
		authMiddleware: am,
	}
	return controller
}
