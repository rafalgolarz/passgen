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
	"github.com/rafalgolarz/passgen/pkg/passwords"
	"github.com/sirupsen/logrus"
)

const (
	ConfigFile = "passgen.toml"
	//PasswortType refers to the section name in the ConfigFile
	PasswordType = "default"
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
	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		log.Debug("Config file " + ConfigFile + " does not exist.")
	} else {
		if _, err := toml.DecodeFile(ConfigFile, &config); err != nil {
			log.Debug("Error loading or parsing the file: " + ConfigFile)
			//to consider: load defaultConfig in case of errors with the file
		} else {
			log.Info("Config file " + ConfigFile + " loaded successfully")
		}
	}
}

// url/cli params that are not passed, will be initialised with default config settings
func initParams(params *passwords.Setting, config passwords.Settings) {

	passwordType := PasswordType
	params.MinLength = config[passwordType].MinLength
	params.MinSpecialCharacters = config[passwordType].MinSpecialCharacters
	params.MinDigits = config[passwordType].MinDigits
	params.MinLowercase = config[passwordType].MinLowercase
	params.MinUppercase = config[passwordType].MinUppercase
	params.Results = config[passwordType].Results
}
