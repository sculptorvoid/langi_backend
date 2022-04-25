package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"net/http"
	"strconv"
)

func (h *Handler) createWord(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dictId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid dictionary id param")
		return
	}

	var input entity.Word
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Word.Create(userId, dictId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllWords(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dictId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid dictionary id param")
		return
	}

	words, err := h.services.Word.GetAll(userId, dictId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, words)
}

func (h *Handler) getWordById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	wordId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid dictionary id param")
		return
	}

	word, err := h.services.Word.GetById(userId, wordId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, word)
}

func (h *Handler) updateWord(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateWordInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Word.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteWord(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	wordId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid dictionary id param")
		return
	}

	err = h.services.Word.Delete(userId, wordId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
