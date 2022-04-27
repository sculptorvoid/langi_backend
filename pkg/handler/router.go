package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sculptorvoid/langi_backend/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			}
		}
		words := api.Group("words")
		{
			words.GET("/:id", h.getWordById)
			words.PUT("/:id", h.updateWord)
			words.DELETE("/:id", h.deleteWord)
		}
	}

	return router
}
