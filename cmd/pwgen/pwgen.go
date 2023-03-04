package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/8ff/diceware"
)

func genRandRange(start int, end int) int {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(end-start+1)))
	if err != nil {
		panic(err)
	}
	return int(num.Int64()) + start
}

func genDiceToken(size int) string {
	dl := diceware.GetWordsMap()
	var token string
	for q := 0; q < size; q++ {
		var rolls string
		for i := 1; i < 6; i++ {
			rolls += fmt.Sprintf("%v", genRandRange(1, 6))
		}
		if q == (size - 1) {
			token += dl[rolls]
		} else {
			token += fmt.Sprintf("%s ", dl[rolls])
		}
	}
	return token
}

func main() {
	// Default length
	length := 6

	// Parse command-line options
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-h" || os.Args[i] == "--help" {
			// Print help message with default length value
			fmt.Printf("NAME:\n  pwgen - Diceware password generator.\n\nUSAGE:\n  %s [OPTIONS]\n\nOPTIONS:\n  -h, --help\t\tShow this help message\n  -l, --length LENGTH\tSpecify the length of the password to generate (default: %d)\n", os.Args[0], length)
			return
		} else if os.Args[i] == "-l" || os.Args[i] == "--length" {
			if i+1 >= len(os.Args) {
				fmt.Printf("ERROR: Length option requires a value\n")
				os.Exit(1)
			}
			var err error
			length, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Printf("ERROR: Invalid length option: %s\n", os.Args[i+1])
				os.Exit(1)
			}
			i++
		} else {
			fmt.Printf("ERROR: Invalid option: %s\n", os.Args[i])
			os.Exit(1)
		}
	}

	q := genDiceToken(length)
	fmt.Printf("%s\n", strings.TrimSpace(q))
}
