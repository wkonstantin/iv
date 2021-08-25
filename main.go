package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	fmt.Println("Start")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "plain/text", []byte("hello world") )
	})
	r.Run()
}
