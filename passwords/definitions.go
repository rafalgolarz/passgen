/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package passwords

const (
	ConfigFile     = "passwords/config.toml"
	AbsoluteMinLen = 8
	//passwortType refers to the section name in the configFile
	PasswordType = "default"
)

type Setting struct {
	MinLength            uint8 `toml:"min_length" form:"min-length"`
	MinSpecialCharacters uint8 `toml:"min_special_characters" form:"min-specials"`
	MinDigits            uint8 `toml:"min_digits" form:"min-digits"`
	MinLowercase         uint8 `toml:"min_lowercase" form:"min-lowers"`
	MinUppercase         uint8 `toml:"min_uppercase" form:"min-uppers"`
	Results              uint8 `toml:"results" form:"res"`
}

type Settings map[string]Setting

type runes struct {
	LowerLetters []rune
	UpperLetters []rune
	Specials     []rune
	Digits       []rune
}

//should match [default] settings in the config file (default: config.toml)
var DefaultConfig = Setting{
	MinLength:            8,
	MinSpecialCharacters: 2,
	MinDigits:            2,
	MinLowercase:         1,
	MinUppercase:         1,
	Results:              1,
}

var AllowedChars = runes{
	LowerLetters: []rune("abcdefghijklmnopqrstuvwxyz"),
	UpperLetters: []rune("ABCDEFGHIHJKLMNOPQRSTUVWXYZ"),
	Specials:     []rune("~!@#$%^&*()_+-=|{}[]\\/';:"),
	Digits:       []rune("0123456789"),
}
