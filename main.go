/*
 * Secure passwords generator
 * sample calls:
	 /v1/passwords returns a password based on default settings
	 /v1/passwords/?min-length=12&min-specials=3&min-digits=3&min-lowers=3&min-uppers=3&res=2
 * @author: rafalgolarz.com
 *
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafalgolarz/passgen/passwords"
	"github.com/sirupsen/logrus"
)

var (
	log    = logrus.New()
	port   string
	config passwords.Settings
)

func init() {
	//set logging level based on GIN_MODE env variable
	setLoggingLevel()
	setAPIListeningPort()
	loadConfigFile()

}

func generatePassword(c *gin.Context) {

	var params passwords.Setting

	initParams(&params, config)
	err := c.BindQuery(&params)
	checkErr(err)

	numberOfResults := int(params.Results)
	passwordsList := make([]string, numberOfResults)
	if checkParams(params) {
		for i := 0; i < numberOfResults; i++ {
			passwordsList[i] = passwords.Generate(params)
		}

		c.JSON(http.StatusOK, gin.H{
			"passwords":      passwordsList,
			"status":         "Success",
			"default_config": passwords.DefaultConfig,
			"applied_config": params})
	} else {
		log.Info("Incorrect password configuration")
		c.JSON(http.StatusNotAcceptable,
			gin.H{
				"passwords":      "",
				"status":         "Incorrect params. Check if params meet default_config",
				"default_config": passwords.DefaultConfig,
				"applied_config": params})
	}

	return
}

func main() {

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/passwords/", generatePassword)
	}
	router.Run(port)
}
