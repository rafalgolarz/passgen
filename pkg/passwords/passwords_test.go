/*
 * Secure passwords generator
 * @author: rafalgolarz.com
 *
 * $ go test -v ./... -bench=./...
 */
package passwords

import "testing"

func BenchmarkGenerator(b *testing.B) {
	var defaultParams, strongParams, superStrong Setting
	defaultParams.MinLength = 8
	defaultParams.MinSpecialCharacters = 2
	defaultParams.MinDigits = 2
	defaultParams.MinLowercase = 1
	defaultParams.MinUppercase = 1

	for i := 0; i < b.N; i++ {
		Generate(defaultParams)
	}

	strongParams.MinLength = 16
	strongParams.MinSpecialCharacters = 4
	strongParams.MinDigits = 4
	strongParams.MinLowercase = 2
	strongParams.MinUppercase = 2

	for i := 0; i < b.N; i++ {
		Generate(strongParams)
	}

	superStrong.MinLength = 255
	superStrong.MinSpecialCharacters = 100
	superStrong.MinDigits = 50
	superStrong.MinLowercase = 20
	superStrong.MinUppercase = 20

	for i := 0; i < b.N; i++ {
		Generate(superStrong)
	}
}
