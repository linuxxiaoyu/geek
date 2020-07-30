package main

import (
	"github.com/gin-gonic/gin"
)

type msg struct {
	Name    string `form:"user_name" binding:"required"`
	Email   string `form:"user_email" binding:"required"`
	Message string `form:"user_message"`
}

func main() {
	r := gin.Default()
	r.POST("/my-handling-form-page", func(c *gin.Context) {
		var m msg
		err := c.ShouldBind(&m)
		if err != nil {
			c.String(400, "Not Found")
			return
		}
		c.JSONP(200, m)
	})
	r.Run(":80")
}
