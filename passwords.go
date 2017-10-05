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

	var charsMix []rune

	minLen := params.MinLength
	minSpecials := params.MinSpecialCharacters
	minDigits := params.MinDigits
	minLowers := params.MinLowercase
	minUppers := params.MinUppercase
	passLen := minSpecials + minDigits + minLowers + minUppers

	// it may happen (and is allowed) that the number of required digits, specials,
	// lowers, uppers is higher than the minimum length, then the new mininum is the sum of lenghts
	log.Info("min required length: ", minLen, ", min length to be generated: ", passLen)

	rand.Seed(time.Now().UTC().UnixNano())

	specials := randChars(int(minSpecials), allowedChars.Specials)
	digits := randChars(int(minDigits), allowedChars.Digits)
	lowerLetters := randChars(int(minLowers), allowedChars.LowerLetters)
	upperLetters := randChars(int(minUppers), allowedChars.UpperLetters)

	var allChars []rune
	allChars = append(allChars, allowedChars.Specials...)
	allChars = append(allChars, allowedChars.Digits...)
	allChars = append(allChars, allowedChars.LowerLetters...)
	allChars = append(allChars, allowedChars.UpperLetters...)

	charsMix = append(charsMix, digits...)
	charsMix = append(charsMix, specials...)
	charsMix = append(charsMix, lowerLetters...)
	charsMix = append(charsMix, upperLetters...)

	log.Info("concatenated (pre-shuffled & pre-reviewed): ", string(charsMix), " [length: ", len(string(charsMix)), "]")
	if minLen > passLen {
		gapSize := int(minLen - passLen)
		gap := make([]rune, gapSize)
		for i := 0; i < gapSize; i++ {
			gap[i] = allChars[rand.Intn(len(allChars))]
		}
		log.Info("gap: ", string(gap), " [length: ", len(string(gap)), "]")
		charsMix = append(charsMix, gap...)
		log.Info("with the gap: ", string(charsMix), " [length: ", len(string(charsMix)), "]")
	}

	log.Info("concatenated (pre shuffled): ", string(charsMix), " [length: ", len(string(charsMix)), "]")

	//shuffle
	for i := range charsMix {
		j := rand.Intn(i + 1)
		charsMix[i], charsMix[j] = charsMix[j], charsMix[i]
	}
	log.Info("final (shuffled) password: ", string(charsMix), " [length: ", len(string(charsMix)), "]")

	return string(charsMix)
}

func randChars(count int, subset []rune) []rune {
	res := make([]rune, count)

	for i := range res {
		res[i] = subset[rand.Intn(len(subset))]
	}

	log.Info("Select ", count, " chars from ", int(len(subset)), " characters: ", string(subset), " RESULT: ", string(res))
	return res
}
