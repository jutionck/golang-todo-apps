package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/domain/dto"
	"github.com/jutionck/golang-todo-apps/usecase"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"net/http"
)

type AuthController struct {
	router *gin.RouterGroup
	uc     usecase.AuthenticationUseCase
}

// Register godoc
// @Summary Register a user
// @Description Register a user
// @Accept json
// @Produce json
// @Tags Auth
// @Param Body body dto.RegisterRequestDto true "Login"
// @Success 201 {object} domain.User
// @Router /auth/register [post]
func (a *AuthController) registerHandler(c *gin.Context) {
	var payload dto.RegisterRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userData := &domain.User{
		Email:    payload.Email,
		Password: payload.Password,
	}
	err := a.uc.Register(userData)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commons.SendCreatedResponse(c, userData, "OK")
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Accept json
// @Produce json
// @Tags Auth
// @Param Body body dto.LoginRequestDto true "Login"
// @Success 201 {object} string
// @Router /auth/login [post]
func (a *AuthController) loginHandler(c *gin.Context) {
	var payload dto.LoginRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		commons.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := a.uc.Login(payload.Email, payload.Password)
	if err != nil {
		commons.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	commons.SendCreatedResponse(c, token, "OK")
}

func (a *AuthController) Route() {
	a.router.POST("/auth/login", a.loginHandler)
	a.router.POST("/auth/register", a.registerHandler)
}

func NewAuthController(r *gin.RouterGroup, uc usecase.AuthenticationUseCase) *AuthController {
	return &AuthController{router: r, uc: uc}
}
