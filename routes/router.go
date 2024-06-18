package routes

import (
	"init_golang/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	api "init_golang/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/user/register", api.UserRegister)
	}
	return r
}
