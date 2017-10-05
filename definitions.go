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
	MinLength         int `toml:"min_length" form:"min-length"`
	SpecialCharacters int `toml:"special_characters" form:"special-chars"`
	Numbers           int `toml:"numbers" form:"numbers"`
	MinLowercase      int `toml:"min_lowercase" form:"min-lower"`
	MinUppercase      int `toml:"min_uppercase" form:"min-upper"`
	Results           int `toml:"results" form:"res"`
}

type settings map[string]setting
