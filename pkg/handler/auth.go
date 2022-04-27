package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"net/http"
)

// @Summary      registration
// @Tags         Authorization
// @Description  create account
// @ID           registration
// @Accept       json
// @Produce      json
// @Param        input    body       entity.User  true  "account info"
// @Success      200      {integer}  registrationResponse
// @Failure      400,404  {object}   errorResponse
// @Failure      500      {object}   errorResponse
// @Failure      default  {object}   errorResponse
// @Router       /auth/registration [post]
func (h *Handler) registration(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, registrationResponse{
		Id: id,
	})
}

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary      login
// @Tags         Authorization
// @Description  login into app
// @ID           login
// @Accept       json
// @Produce      json
// @Param        input    body      loginInput  true  "credentials"
// @Success      200      {object}  loginResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input loginInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		Token: token,
	})
}
