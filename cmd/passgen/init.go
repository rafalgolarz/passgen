/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/rafalgolarz/passgen/pkg/passwords"
	"github.com/sirupsen/logrus"
)

const (
	DefaultConfigFile = "passgen.toml"
	//PasswortType refers to the section name in the ConfigFile
	PasswordType = "default"
)

func setLoggingLevel() {
	//for production log only: Error, Fatal and Panic.
	//for non-production log: Debug, Info, Warning, Error, Fatal and Panic
	log.SetLevel(logrus.ErrorLevel)
}

func loadConfigFile(cmdConfigPath string) {

	if cmdConfigPath != DefaultConfigFile {
		if _, err := os.Stat(cmdConfigPath); os.IsNotExist(err) {
			log.Error("Config file " + cmdConfigPath + " does not exist.")
		}
	} else if _, err := os.Stat(DefaultConfigFile); os.IsNotExist(err) {
		log.Debug("Default config file " + DefaultConfigFile + " does not exist.")
	}

	if _, err := toml.DecodeFile(cmdConfigPath, &config); err != nil {
		log.Debug("Error loading or parsing the file: " + cmdConfigPath)
	} else {
		log.Info("Config file " + cmdConfigPath + " loaded successfully")
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
