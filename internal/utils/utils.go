package utils

import (
	"math"
	"unicode"
	"unicode/utf8"


)

func CapitalizeFirstLetter(s string) string {
    if len(s) == 0 {
        return ""
    }
    r, size := utf8.DecodeRuneInString(s)
    return string(unicode.ToUpper(r)) + s[size:]
}

func FirstLetterFromString(s string) string {
	r, _ := utf8.DecodeRuneInString(s)
	return string(r)
}

func RoundFloat(f float64, decimalPlaces int) float64 {

    div := math.Pow(10.0, float64(decimalPlaces))

    return math.Round(f * div) / div

}

