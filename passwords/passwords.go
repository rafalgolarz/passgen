/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package passwords

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Generate returns a password built based on passed requirements
func Generate(params Setting) string {

	minLen := params.MinLength
	minSpecials := params.MinSpecialCharacters
	minDigits := params.MinDigits
	minLowers := params.MinLowercase
	minUppers := params.MinUppercase
	paramsLen := minSpecials + minDigits + minLowers + minUppers

	// it may happen (and is allowed) that the number of required digits, specials,
	// lowers, uppers is higher than the minimum length, then the new mininum is the sum of lenghts
	log.Info("min required length: ", minLen, ", lenth of the required params: ", paramsLen)

	rand.Seed(time.Now().UTC().UnixNano())

	specials := randChars(int(minSpecials), AllowedChars.Specials)
	digits := randChars(int(minDigits), AllowedChars.Digits)
	lowerLetters := randChars(int(minLowers), AllowedChars.LowerLetters)
	upperLetters := randChars(int(minUppers), AllowedChars.UpperLetters)

	var allChars []rune
	allChars = append(allChars, AllowedChars.Specials...)
	allChars = append(allChars, AllowedChars.Digits...)
	allChars = append(allChars, AllowedChars.LowerLetters...)
	allChars = append(allChars, AllowedChars.UpperLetters...)

	var charsMix []rune
	charsMix = append(charsMix, digits...)
	charsMix = append(charsMix, specials...)
	charsMix = append(charsMix, lowerLetters...)
	charsMix = append(charsMix, upperLetters...)

	log.Info("concatenated (pre-shuffled & pre-reviewed): ", string(charsMix), " [length: ", len(string(charsMix)), "]")

	if minLen > paramsLen {
		gapSize := int(minLen - paramsLen)
		gap := make([]rune, gapSize)

		for i := 0; i < gapSize; i++ {
			gap[i] = allChars[rand.Intn(len(allChars))]
		}

		log.Info("gap: ", string(gap), " [length: ", len(string(gap)), "]")
		charsMix = append(charsMix, gap...)
		log.Info("with the gap: ", string(charsMix), " [length: ", len(string(charsMix)), "]")
	}

	log.Info("concatenated (pre-shuffled): ", string(charsMix), " [length: ", len(string(charsMix)), "]")

	//shuffle
	for i := range charsMix {
		j := rand.Intn(i + 1)
		charsMix[i], charsMix[j] = charsMix[j], charsMix[i]
	}
	log.Info("final (shuffled) password: ", string(charsMix), " [length: ", len(string(charsMix)), "]")

	return string(charsMix)
}

// randCharts return a random combination of charts from []rune (a group of characters)
// subsetLen defines the length of returned subset
func randChars(subsetLen int, characters []rune) []rune {
	res := make([]rune, subsetLen)

	for i := range res {
		res[i] = characters[rand.Intn(len(characters))]
	}

	log.Info("Select ", subsetLen, " chars from ", int(len(characters)), " characters: ", string(characters), " RESULT: ", string(res))
	return res
}
