package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"net/http"
	"strconv"
)

// @Summary      createDictionary
// @Security     ApiKeyAuth
// @Tags         Dictionary
// @Description  creates an empty dictionary
// @ID           createDictionary
// @Accept       json
// @Produce      json
// @Param        input    body      entity.Dictionary  true  "credentials"
// @Success      200      {object}  createDictionaryResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict [post]
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
	c.JSON(http.StatusOK, createDictionaryResponse{
		Id: id,
	})
}

// @Summary      getAllDictionaries
// @Security     ApiKeyAuth
// @Tags         Dictionary
// @Description  return all dictionaries
// @ID           getAllDictionaries
// @Accept       json
// @Produce      json
// @Success      200      {string}  []entity.Dictionary
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict [get]
func (h *Handler) getAllDictionaries(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dictionaries, err := h.services.Dictionary.GetAllDictionaries(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDictionariesResponse{
		Data: dictionaries,
	})
}

// @Summary      getDictionaryById
// @Security     ApiKeyAuth
// @Tags         Dictionary
// @Description  return dictionary by id
// @ID           getDictionaryById
// @Accept       json
// @Produce      json
// @Param        id       body      int  true  "dictionary id"
// @Success      200      {object}  getDictionaryByIdResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict/:id [get]
func (h *Handler) getDictionaryById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	dictionary, err := h.services.Dictionary.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getDictionaryByIdResponse{
		dictionary,
	})
}

// @Summary      updateDictionary
// @Security     ApiKeyAuth
// @Tags         Dictionary
// @Description  update dictionary by id
// @ID           updateDictionary
// @Accept       json
// @Produce      json
// @Param        id       body      int  true  "dictionary id"
// @Success      200      {object}  updateDictionaryResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict/:id [put]
func (h *Handler) updateDictionary(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateDictionaryInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Dictionary.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateDictionaryResponse{
		Id: id,
	})
}

// @Summary      deleteDictionary
// @Security     ApiKeyAuth
// @Tags         Dictionary
// @Description  delete dictionary by id
// @ID           deleteDictionary
// @Accept       json
// @Produce      json
// @Param        id       body      int  true  "dictionary id"
// @Success      200      {object}  deleteDictionaryResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/dict/:id [delete]
func (h *Handler) deleteDictionary(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Dictionary.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteDictionaryResponse{
		Id: id,
	})
}
