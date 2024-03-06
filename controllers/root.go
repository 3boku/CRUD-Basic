package controllers

import (
	"crud-server/repositories"
	"crud-server/services"

	"github.com/gin-gonic/gin"
)

func NewController(port string) {
	r := gin.New()
	r.Use(gin.Logger())

	mongo := repositories.MongoInit();

	r.GET("/get-feed", func (c *gin.Context)  {
		services.GetFeed(c)
	})
	r.POST("/post-feed", func (c *gin.Context) {
		services.Post(c, &mongo)
	})

	err := r.Run(port)

	if err != nil {
		panic(err.Error())
	}
}