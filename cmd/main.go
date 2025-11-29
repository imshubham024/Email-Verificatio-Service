package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imshubham024/go-emailVerify/api"
)

func main() {
	router := gin.Default()
	// initialize the config
	app := api.Config{Router: router}
	app.Routes()
	router.Run(":8082")
}
