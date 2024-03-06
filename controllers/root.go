package controllers

import "github.com/gin-gonic/gin"

func NewController(port string) {
	r := gin.New()

	r.Use(gin.Logger())

	err := r.Run(port)
	if err != nil {
		panic(err.Error())
	}
}