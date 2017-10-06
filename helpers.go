/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package main

import (
	"github.com/rafalgolarz/passgen/passwords"
)

// check if a combination of values makes sense
func checkParams(params passwords.Setting) bool {

	minLen := params.MinLength
	minSpecials := params.MinSpecialCharacters
	minDigits := params.MinDigits
	minLowers := params.MinLowercase
	minUppers := params.MinUppercase
	passwordType := passwords.PasswordType

	if minLen < passwords.AbsoluteMinLen ||
		minDigits < config[passwordType].MinDigits ||
		minLowers < config[passwordType].MinLowercase ||
		minUppers < config[passwordType].MinUppercase ||
		minSpecials < config[passwordType].MinSpecialCharacters {
		return false
	}

	return true
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
