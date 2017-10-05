/*
 * Secure passwords generator
 * sample calls:
	 /v1/passwords returns a password based on default settings
	 /v1/passwords/?min-length=12&min-specials=3&min-digits=3&min-lowers=3&min-uppers=3&res=2
 * @author: rafal@rafalgolarz.com
 *
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	log    = logrus.New()
	port   string
	config settings
)

func init() {
	//set logging level based on GIN_MODE env variable
	setLoggingLevel()
	setAPIListeningPort()
	loadConfigFile()

}

func generatePassword(c *gin.Context) {

	var params setting
	initParams(&params, config)
	err := c.BindQuery(&params)
	checkErr(err)

	//the whole logic will go here:
	//check if params make sense (len>sum of allowed chars, etc)
	//randomiser goes to another file

	if checkParams(params) {
		c.JSON(http.StatusOK, gin.H{"password configuration": params})
	} else {
		c.JSON(http.StatusNotAcceptable,
			gin.H{"status": "Incorrect password configuration",
				"Minimum acceptable values of params": defaultConfig})
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
