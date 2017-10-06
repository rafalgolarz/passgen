/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 */
package passwords

const (
	ConfigFile     = "passwords/config.toml"
	AbsoluteMinLen = 8
	//PasswortType refers to the section name in the ConfigFile
	PasswordType = "default"
)

// UnsignedInt type used to set limits of allowed vales for url params
// Have fun experimenting
// uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255. You won't need more...
// uint16. Range: 0 through 65535. works really well on my machine.
// uint32. Range: 0 through 4294967295 ...but this one may kill your memory

type UnsignedInt uint8

type Setting struct {
	MinLength            UnsignedInt `toml:"min_length" form:"min-length"`
	MinSpecialCharacters UnsignedInt `toml:"min_special_characters" form:"min-specials"`
	MinDigits            UnsignedInt `toml:"min_digits" form:"min-digits"`
	MinLowercase         UnsignedInt `toml:"min_lowercase" form:"min-lowers"`
	MinUppercase         UnsignedInt `toml:"min_uppercase" form:"min-uppers"`
	Results              UnsignedInt `toml:"results" form:"res"`
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
