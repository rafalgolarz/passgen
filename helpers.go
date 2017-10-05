/*
 * Secure passwords generator
 * @author: rafal@rafalgolarz.com
 *
 */
package main

func checkParams(params setting) bool {

	// minLen := params.MinLength
	// specials := params.SpecialCharacters
	// numbers := params.Numbers
	// lowers := params.MinLowercase
	// uppers := params.MinUppercase
	return true
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
