/*
 * Secure passwords generator
 * @author: rafal@rafalgolarz.com
 *
 */
package main

const (
	configFile     = "config.toml"
	defaultEnvPort = ":8080"
)

var defaultConfig settings

type setting struct {
	MinLength         uint8 `toml:"min_length" form:"min-length"`
	SpecialCharacters uint8 `toml:"special_characters" form:"special-chars"`
	Numbers           uint8 `toml:"numbers" form:"numbers"`
	MinLowercase      uint8 `toml:"min_lowercase" form:"min-lower"`
	MinUppercase      uint8 `toml:"min_uppercase" form:"min-upper"`
	Results           uint8 `toml:"results" form:"res"`
}

type settings map[string]setting
