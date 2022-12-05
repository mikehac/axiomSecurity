package main

import (
	"axiomsecurity/manifestlib"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	manifestlib.InitComboBox()
	message := manifestlib.Hello("Mike")
	fmt.Println(message)

	//Creating a web server
	router := gin.Default()
	router.GET("/form/:name", getFormManifest)
	router.GET("/values/:source", getValues)
	router.POST("/form/:name", postForm)

	router.Run("localhost:3000")

}
