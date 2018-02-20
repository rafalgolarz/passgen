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

const (
	AbsoluteMinLen = 8
)

var (
	// compile passing -ldflags "-X main.Build <build sha1>"
	build  string
	config Settings
)

// UnsignedInt type used to set limits of allowed vales for url params
// Have fun experimenting
// uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255. You won't need more...
// uint16. Range: 0 through 65535. works really well on my machine.
// uint32. Range: 0 through 4294967295 ...but this one may kill your memory
type UnsignedInt uint8

//CharsSet definition
//rune as actualy an alias for the type int32
type CharsSet []rune

//Setting represents a structure used by both config file and url params
type Setting struct {
	MinLength            UnsignedInt `toml:"min_length" form:"min-length"`
	MinSpecialCharacters UnsignedInt `toml:"min_special_characters" form:"min-specials"`
	MinDigits            UnsignedInt `toml:"min_digits" form:"min-digits"`
	MinLowercase         UnsignedInt `toml:"min_lowercase" form:"min-lowers"`
	MinUppercase         UnsignedInt `toml:"min_uppercase" form:"min-uppers"`
	Results              UnsignedInt `toml:"results" form:"res"`
}

//Settings map a section name to settings
type Settings map[string]Setting

type runes struct {
	LowerLetters CharsSet
	UpperLetters CharsSet
	Specials     CharsSet
	Digits       CharsSet
}

//DefaultConfig should match [default] settings in the config file (default: configs/default.toml)
var DefaultConfig = Setting{
	MinLength:            8,
	MinSpecialCharacters: 2,
	MinDigits:            2,
	MinLowercase:         1,
	MinUppercase:         1,
	Results:              1,
}

//AllowedChars define acceptable chars used to generate passwords
var AllowedChars = runes{
	LowerLetters: CharsSet("abcdefghijklmnopqrstuvwxyz"),
	UpperLetters: CharsSet("ABCDEFGHIHJKLMNOPQRSTUVWXYZ"),
	Specials:     CharsSet("~!@#$%^&*()_+-=|{}[]\\/';:"),
	Digits:       CharsSet("0123456789"),
}

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

	var allChars CharsSet
	allChars = append(allChars, AllowedChars.Specials...)
	allChars = append(allChars, AllowedChars.Digits...)
	allChars = append(allChars, AllowedChars.LowerLetters...)
	allChars = append(allChars, AllowedChars.UpperLetters...)

	var charsMix CharsSet
	charsMix = append(charsMix, digits...)
	charsMix = append(charsMix, specials...)
	charsMix = append(charsMix, lowerLetters...)
	charsMix = append(charsMix, upperLetters...)

	log.Info("concatenated (pre-shuffled & pre-reviewed): ", string(charsMix), " [length: ", len(string(charsMix)), "]")

	if minLen > paramsLen {
		gapSize := int(minLen - paramsLen)
		gap := make(CharsSet, gapSize)

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

// randCharts return a random combination of charts from CharsSet (a group of characters)
// subsetLen defines the length of returned subset
func randChars(subsetLen int, characters CharsSet) CharsSet {
	res := make(CharsSet, subsetLen)

	for i := range res {
		res[i] = characters[rand.Intn(len(characters))]
	}

	log.Info("Select ", subsetLen, " chars from ", int(len(characters)), " characters: ", string(characters), " RESULT: ", string(res))
	return res
}

// check if a combination of values makes sense
func CheckParams(params Setting, passwordType string) bool {

	minLen := params.MinLength
	minSpecials := params.MinSpecialCharacters
	minDigits := params.MinDigits
	minLowers := params.MinLowercase
	minUppers := params.MinUppercase

	if minLen < AbsoluteMinLen ||
		minDigits < config[passwordType].MinDigits ||
		minLowers < config[passwordType].MinLowercase ||
		minUppers < config[passwordType].MinUppercase ||
		minSpecials < config[passwordType].MinSpecialCharacters {
		return false
	}

	return true
}

func CheckErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
