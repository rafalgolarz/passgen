/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/rafalgolarz/passgen/passwords"
	"github.com/sirupsen/logrus"
)

func setLoggingLevel() {
	if gin.Mode() == "release" {
		//for production log only: Error, Fatal and Panic.
		log.SetLevel(logrus.ErrorLevel)
	} else {
		//for non-production log: Debug, Info, Warning, Error, Fatal and Panic
		log.SetLevel(logrus.DebugLevel)
	}
}

//get the default port from environmental variables
func setAPIListeningPort() {
	port = os.Getenv("DEFAULT_API_PORT")
	if port == "" {
		port = ":8080"
	}
}

func loadConfigFile() {
	if _, err := os.Stat(passwords.ConfigFile); os.IsNotExist(err) {
		log.Error("Config file " + passwords.ConfigFile + " does not exist.")
	} else {
		if _, err := toml.DecodeFile(passwords.ConfigFile, &config); err != nil {
			log.Error("Error parsing " + passwords.ConfigFile)
			//to consider: load defaultConfig in case of errors with the file
		} else {
			log.Info("Config file " + passwords.ConfigFile + " loaded successfully")
		}
	}
}

// url params that are not passed, will be initialised with default config settings
func initParams(params *passwords.Setting, config passwords.Settings) {

	passwordType := passwords.PasswordType
	params.MinLength = config[passwordType].MinLength
	params.MinSpecialCharacters = config[passwordType].MinSpecialCharacters
	params.MinDigits = config[passwordType].MinDigits
	params.MinLowercase = config[passwordType].MinLowercase
	params.MinUppercase = config[passwordType].MinUppercase
	params.Results = config[passwordType].Results
}
