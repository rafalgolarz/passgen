/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/rafalgolarz/passgen/pkg/passwords"
	"github.com/sirupsen/logrus"
)

var (
	log    = logrus.New()
	config passwords.Settings
	wg     sync.WaitGroup
)

func generatePassword(params passwords.Setting) passwords.PasswordsList {
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
	} else {
		log.Info("Incorrect password configuration")
	}

	return ps
}

func main() {

	var cmdParams passwords.Setting

	minLength := flag.Uint("min-length", 8, "minimum length of passwords (max. 255)")
	minSpecialCharacters := flag.Uint("min-specials", 2, "minimum number of special characters (max. 255)")
	minDigits := flag.Uint("min-digits", 1, "minimum number of digits (max. 255)")
	minLowercase := flag.Uint("min-lowers", 1, "minimum number of lower case letters (max. 255)")
	minUppercase := flag.Uint("min-uppers", 1, "minimum number of upper case letters (max. 255)")
	results := flag.Uint("res", 1, "number of passwords to generate")
	verbose := flag.Bool("verbose", false, "run in verbose mode (default false)")
	configFile := flag.String("config-file", "passgen.toml", "path to the custom configuration file. must be in the directory where passgen executes.")

	flag.Parse()

	loadConfigFile(*configFile)

	cmdParams.MinLength = passwords.UnsignedInt(*minLength)
	cmdParams.MinSpecialCharacters = passwords.UnsignedInt(*minSpecialCharacters)
	cmdParams.MinDigits = passwords.UnsignedInt(*minDigits)
	cmdParams.MinLowercase = passwords.UnsignedInt(*minLowercase)
	cmdParams.MinUppercase = passwords.UnsignedInt(*minUppercase)
	cmdParams.Results = passwords.MAX_RESULTS(*results)
	passwordsList := generatePassword(cmdParams)

	start := time.Now()
	//casting type will set the value to 0 if not in range
	res := passwords.MAX_RESULTS(*results)

	if res != 0 {
		for _, p := range passwordsList {
			fmt.Println(p)
		}

		if *verbose {
			fmt.Println("Max. possible results: ", passwords.MaxResults())
			fmt.Println("Generated ", len(passwordsList), " passwords in ", time.Since(start))
		}
	} else {
		fmt.Println("The maximum number of results can be set to ", passwords.MaxResults())
	}

}
