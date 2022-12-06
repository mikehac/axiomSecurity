package main

import (
	"axiomsecurity/manifestlib"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadDataFromJsonFiles()

	//Creating a web server
	router := gin.Default()
	router.GET("/form/:name", getFormManifest)
	router.GET("/values/:source", getValues)
	router.POST("/form/:name", postForm)

	router.Run("localhost:3000")

}

func LoadDataFromJsonFiles() {
	// accounts := manifestlib.Accounts{}
	// accounts.LoadData()
	// dbInstances := manifestlib.DatabaseInstances{}
	// dbInstances.LoadData()
	// databaseManifest := manifestlib.Manifest{}
	manifestlib.LoadSourceValues()
	manifestlib.LoadManifest()
}
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
