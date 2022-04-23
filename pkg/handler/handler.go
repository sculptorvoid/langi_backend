package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculptorvoid/langi_backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/registration", h.registration)
	}

	api := router.Group("/api", h.userIdentity)
	{
		dict := api.Group("/dict")
		{
			dict.POST("/", h.createDictionary)
			dict.GET("/", h.getAllDictionaries)
			dict.GET("/:id", h.getDictionaryById)
			dict.PUT("/:id", h.updateDictionary)
			dict.DELETE("/:id", h.deleteDictionary)

			words := dict.Group(":id/words")
			{
				words.POST("/", h.createWord)
				words.GET("/", h.getAllWords)
				words.GET("/:word_id", h.getWordById)
				words.PUT("/:word_id", h.updateWord)
				words.DELETE("/:word_id", h.deleteWord)
			}
		}
	}

	return router
}
