/*
 * Secure passwords generator
 * @author: rafal@rafalgolarz.com
 *
 */
package main

import (
	"math/rand"
	"time"
)

func generate(params setting) string {

	minLen := params.MinLength
	minSpecials := params.MinSpecialCharacters
	minDigits := params.MinDigits
	minLowers := params.MinLowercase
	minUppers := params.MinUppercase

	rand.Seed(time.Now().UTC().UnixNano())
	charsMix := make([]rune, minLen)

	specials := randChars(int(minSpecials), allowedChars.Specials)
	digits := randChars(int(minDigits), allowedChars.Digits)
	lowerLetters := randChars(int(minLowers), allowedChars.LowerLetters)
	upperLetters := randChars(int(minUppers), allowedChars.UpperLetters)

	charsMix = append(charsMix, digits...)
	charsMix = append(charsMix, specials...)
	charsMix = append(charsMix, lowerLetters...)
	charsMix = append(charsMix, upperLetters...)

	return string(charsMix)
}

func randChars(count int, subset []rune) []rune {
	res := make([]rune, count)

	for i := range res {
		res[i] = subset[rand.Intn(len(subset))]
	}
	return res
}
