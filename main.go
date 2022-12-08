package main

import (
	"axiomsecurity/manifestlib"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadDataFromJsonFiles()

	//Creating a web server
	initRoutes()
}

func initRoutes() {
	router := gin.Default()
	router.GET("/form/:name", getFormManifest)
	router.GET("/values/:source", getValues)
	router.POST("/form/:name", postForm)

	router.Run("localhost:3000")
}

func LoadDataFromJsonFiles() {
	manifestlib.LoadSourceValues()
	manifestlib.BuildManifests()
}
func getFormManifest(c *gin.Context) {
	formName := c.Param("name")
	if currentManifest, ok := manifestlib.ManifestsMap[formName]; ok {
		c.JSON(http.StatusOK, currentManifest)
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": string("No manifest named '" + formName + "' was found")})
	}
}

func getValues(c *gin.Context) {
	source := c.Param("source")
	if currentSource, responseStatus := manifestlib.SourceValueMap[source]; responseStatus {
		c.JSON(http.StatusOK, currentSource)
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": string("No source named '" + source + "' was found")})
	}
}

func postForm(c *gin.Context) {
	formName := c.Param("name")
	formDataToValidate, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": string("Error:" + err.Error())})
	}

	if manifestlib.ValidateTypeAndValues(formDataToValidate, formName) {
		c.JSON(http.StatusOK, "success")
	} else {
		c.JSON(http.StatusOK, "fail")
	}
}
