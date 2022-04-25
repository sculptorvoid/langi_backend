package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"net/http"
)

func (h *Handler) createDictionary(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input entity.Dictionary
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Dictionary.CreateDictionary(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllDictionaries(c *gin.Context) {

}

func (h *Handler) getDictionaryById(c *gin.Context) {

}

func (h *Handler) updateDictionary(c *gin.Context) {

}

func (h *Handler) deleteDictionary(c *gin.Context) {

}
