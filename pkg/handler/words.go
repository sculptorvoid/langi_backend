package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"net/http"
	"strconv"
)

// @Summary      createWord
// @Security     ApiKeyAuth
// @Tags         Words
// @Description  creates a word
// @ID           createWord
// @Accept       json
// @Produce      json
// @Param        input    body      entity.Word  true  "word and translation"
// @Success      200      {object}  createWordResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict/:id/words [post]
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

	wordId, err := h.services.Word.Create(userId, dictId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createWordResponse{
		Id: wordId,
	})
}

// @Summary      getAllWords
// @Security     ApiKeyAuth
// @Tags         Words
// @Description  return all words in dictionary
// @ID           getAllWords
// @Accept       json
// @Produce      json
// @Success      200      {object}  getAllWordsResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict/:id/words [get]
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

	c.JSON(http.StatusOK, getAllWordsResponse{
		words,
	})
}

// @Summary      getWordById
// @Security     ApiKeyAuth
// @Tags         Words
// @Description  return word by id from dictionary
// @ID           getWordById
// @Accept       json
// @Produce      json
// @Param        id       body      int  true  "word id"
// @Success      200      {object}  getWordByIdResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/words/:id [get]
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

	c.JSON(http.StatusOK, getWordByIdResponse{
		word,
	})
}

// @Summary      updateWord
// @Security     ApiKeyAuth
// @Tags         Words
// @Description  update word by id from dictionary
// @ID           updateWord
// @Accept       json
// @Produce      json
// @Param        id       body      int  true  "word id"
// @Success      200      {object}  updateWordResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/words/:id [put]
func (h *Handler) updateWord(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	wordId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateWordInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Word.Update(userId, wordId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateWordResponse{
		Id: wordId,
	})
}

// @Summary      deleteWord
// @Security     ApiKeyAuth
// @Tags         Words
// @Description  delete word by id from dictionary
// @ID           deleteWord
// @Accept       json
// @Produce      json
// @Param        id       body      int  true  "word id"
// @Success      200      {object}  deleteWordResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/words/:id [delete]
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

	c.JSON(http.StatusOK, deleteWordResponse{
		Id: wordId,
	})
}
