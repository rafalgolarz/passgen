/*
 * Secure passwords generator
 * @author: rafal@rafalgolarz.com
 *
 */
package main

// check if a combination of values makes sense
func checkParams(params setting) bool {

	minLen := params.MinLength
	minSpecials := params.MinSpecialCharacters
	minDigits := params.MinDigits
	minLowers := params.MinLowercase
	minUppers := params.MinUppercase

	if minLen < absoluteMinLen ||
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
