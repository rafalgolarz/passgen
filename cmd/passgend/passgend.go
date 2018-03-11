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
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rafalgolarz/passgen/pkg/passwords"
	"github.com/sirupsen/logrus"
)

var (
	log    = logrus.New()
	port   string
	config passwords.Settings
	wg     sync.WaitGroup
)

func init() {
	//set logging level based on GIN_MODE env variable
	setLoggingLevel()
	setAPIListeningPort()
	loadConfigFile()

}

func generatePassword(c *gin.Context) {

	var params passwords.Setting

	err := c.BindQuery(&params)
	passwords.CheckErr(err)

	numberOfResults := int(params.Results)
	ps := make(passwords.PasswordsList, numberOfResults)

	if passwords.CheckParams(params, PasswordType) {
		wg.Add(numberOfResults)
		var mutex = &sync.Mutex{}

		for i := 0; i < numberOfResults; i++ {
			go func(i int, ps passwords.PasswordsList) {
				mutex.Lock()
				ps[i] = passwords.Generate(params, &wg)
				mutex.Unlock()
			}(i, ps)
		}

		wg.Wait()

		c.JSON(http.StatusOK, gin.H{
			"passwords":      ps,
			"status":         "Success",
			"default_config": passwords.DefaultConfig,
			"applied_config": params})
	} else {
		log.Info("Incorrect password configuration")
		c.JSON(http.StatusNotAcceptable,
			gin.H{
				"passwords":      []string{},
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
