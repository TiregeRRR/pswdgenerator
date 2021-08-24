package pswdgenerator

import (
	"math/rand"
	"strings"
	"time"
)

type generatorAlph struct {
	alphabet          []string
	numbers           []string
	specialCharacters []string
	symbols           []string
}

var presets generatorAlph

func (g generatorAlph) presetsInit() {
	g.alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	g.numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	g.specialCharacters = []string{"{", "}", "[", "]", "(", ")", "/", `\`, `'`, `"`, `~`, `;`, `:`, `.`, `<`, `>`}
	g.symbols = []string{"@", "#", "$", "%"}
}

// GetPswdMsk returns a string password, msk support C, c - for big and low case letter, N - for numbers,
// S - Special characters, T - Symbols. If msk is empty string call a GetPswdLen with length = 10,
func GetPswdMsk(msk string, special bool) string {
	rand.Seed(time.Now().Unix())
	if msk == "" {
		return GetPswdLen(10, special)
	} else {
		presets.presetsInit()
		pswd := ""
		for _, v := range msk {
			switch v {
			case 'C':
				pswd += presets.alphabet[rand.Intn(26)]
			case 'c':
				pswd += strings.ToLower(presets.alphabet[rand.Intn(26)])
			case 'N':
				pswd += presets.numbers[rand.Intn(10)]
			case 'n':
				pswd += presets.numbers[rand.Intn(10)]
			case 'S':
				pswd += presets.specialCharacters[rand.Intn(16)]
			case 's':
				pswd += presets.specialCharacters[rand.Intn(16)]
			case 'T':
				pswd += presets.symbols[rand.Intn(4)]
			case 't':
				pswd += presets.symbols[rand.Intn(4)]
			}
		}
		return pswd
	}
}

// GetPswdLen returns a string password, you need to pass length of password
func GetPswdLen(len int, special bool) string {
	rand.Seed(time.Now().Unix())
	pswdMsk := ""
	if special {
		for i := 0; i < len; i++ {
			switch rand.Intn(4) {
			case 0:
				pswdMsk += "C"
			case 1:
				pswdMsk += "c"
			case 2:
				pswdMsk += "N"
			case 3:
				pswdMsk += "S"
			}
		}
	}
	return GetPswdMsk(pswdMsk, special)
}
