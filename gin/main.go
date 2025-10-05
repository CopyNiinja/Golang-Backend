package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	 
	// Create a Gin router with default middleware (logger and recovery)
	r:=gin.Default()

	//endpoints
	r.GET("/ping",func(c *gin.Context) {
      //JSON response
	  c.JSON(http.StatusOK,gin.H{
		"message":"hello world",
	  })
	})

	//port
	var port  = 8081;
	flag.IntVar(&port,"port",port,"server port number");
    //parsing the flag
	flag.Parse()

	//run and listening the server
	r.Run(":"+strconv.Itoa(port));
}