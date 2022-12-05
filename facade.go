package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getFormManifest(c *gin.Context) {
	formName := c.Param("name")
	fmt.Println(formName)
}

func getValues(c *gin.Context) {
	source := c.Param("source")
	fmt.Println(source)
}

func postForm(c *gin.Context) {
	formName := c.Param("name")
	fmt.Println(formName)
}
