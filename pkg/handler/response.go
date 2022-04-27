package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sirupsen/logrus"
)

// --------- Authorization responses ---------

type registrationResponse struct {
	Id int `json:"id"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// --------- Dictionary responses ---------

type createDictionaryResponse struct {
	Id int `json:"id"`
}

type getAllDictionariesResponse struct {
	Data []entity.Dictionary `json:"data"`
}

type getDictionaryByIdResponse struct {
	Data entity.Dictionary `json:"data"`
}

type updateDictionaryResponse struct {
	Id int `json:"id"`
}

type deleteDictionaryResponse struct {
	Id int `json:"id"`
}

// --------- Word responses ---------

type createWordResponse struct {
	Id int `json:"id"`
}

type getAllWordsResponse struct {
	Data []entity.Word `json:"data"`
}

type getWordByIdResponse struct {
	Data entity.Word `json:"data"`
}

type updateWordResponse struct {
	Id int `json:"id"`
}

type deleteWordResponse struct {
	Id int `json:"id"`
}

// --------- Error responses ---------

type errorResponse struct {
	Message string `json:"error"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
