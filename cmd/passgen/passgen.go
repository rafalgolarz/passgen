/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package main

import (
	"flag"
	"fmt"

	"github.com/rafalgolarz/passgen/pkg/passwords"
	"github.com/sirupsen/logrus"
)

var (
	log    = logrus.New()
	config passwords.Settings
)

func init() {
	setLoggingLevel()
}

func generatePassword(params passwords.Setting) []string {

	numberOfResults := int(params.Results)
	passwordsList := make([]string, numberOfResults)
	if passwords.CheckParams(params, PasswordType) {
		for i := 0; i < numberOfResults; i++ {
			passwordsList[i] = passwords.Generate(params)
		}
	} else {
		log.Info("Incorrect password configuration")
	}

	return passwordsList
}

func main() {

	var cmdParams passwords.Setting

	minLength := flag.Uint("min-length", 8, "minimum length of passwords")
	minSpecialCharacters := flag.Uint("min-specials", 2, "minimum number of special characters")
	minDigits := flag.Uint("min-digits", 1, "minimum number of digits")
	minLowercase := flag.Uint("min-lowers", 1, "minimum number of lower case letters")
	minUppercase := flag.Uint("min-uppers", 1, "minimum number of upper case letters")
	results := flag.Uint("res", 1, "number of passwords to generate")
	configFile := flag.String("config-file", "passgen.toml", "path to the custom configuration file. must be in the directory where passgen executes.")

	flag.Parse()

	loadConfigFile(*configFile)

	cmdParams.MinLength = passwords.UnsignedInt(*minLength)
	cmdParams.MinSpecialCharacters = passwords.UnsignedInt(*minSpecialCharacters)
	cmdParams.MinDigits = passwords.UnsignedInt(*minDigits)
	cmdParams.MinLowercase = passwords.UnsignedInt(*minLowercase)
	cmdParams.MinUppercase = passwords.UnsignedInt(*minUppercase)
	cmdParams.Results = passwords.UnsignedInt(*results)

	passwords := generatePassword(cmdParams)

	for _, passw := range passwords {
		fmt.Println(passw)
	}
}
