/*
 * Secure passwords generator
 * @author: rafal@rafalgolarz.com
 *
 */
package main

const (
	configFile     = "config.toml"
	defaultEnvPort = ":8080"
	absoluteMinLen = 8
	//passwortType refers to the section name in the configFile
	passwordType = "default"
)

type setting struct {
	MinLength            uint8 `toml:"min_length" form:"min-length"`
	MinSpecialCharacters uint8 `toml:"min_special_characters" form:"min-specials"`
	MinDigits            uint8 `toml:"min_digits" form:"min-digits"`
	MinLowercase         uint8 `toml:"min_lowercase" form:"min-lowers"`
	MinUppercase         uint8 `toml:"min_uppercase" form:"min-uppers"`
	Results              uint8 `toml:"results" form:"res"`
}

type settings map[string]setting

type runes struct {
	LowerLetters []rune
	UpperLetters []rune
	Specials     []rune
	Digits       []rune
}

//should match [default] settings in the config file (default: config.toml)
var defaultConfig = setting{
	MinLength:            8,
	MinSpecialCharacters: 2,
	MinDigits:            2,
	MinLowercase:         1,
	MinUppercase:         1,
	Results:              1,
}

var allowedChars = runes{
	LowerLetters: []rune("abcdefghijklmnopqrstuvwxyz"),
	UpperLetters: []rune("ABCDEFGHIHJKLMNOPQRSTUVWXYZ"),
	Specials:     []rune("~!@#$%^&*()_+-=|{}[]\\/';:"),
	Digits:       []rune("01234567890"),
}
