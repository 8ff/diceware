![logo](media/dice.svg)

# Diceware

[![Go Report Card](https://goreportcard.com/badge/github.com/8ff/diceware)](https://goreportcard.com/report/github.com/8ff/diceware)
[![GoDoc](https://godoc.org/github.com/8ff/diceware?status.svg)](https://godoc.org/github.com/8ff/diceware)
[![License](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://raw.githubusercontent.com/8ff/diceware/master/LICENSE)

Diceware is a technique for generating strong, memorable passwords using dice to select words from a list. The Diceware method is based on the idea that each word in the list corresponds to a unique combination of five dice rolls, making it difficult for attackers to guess the password.

## Diceware
The `diceware` library provided in this project allows you to generate diceware passwords using Go. The library includes functions for generating a list of words, a randomised list of words, and a map of words that corresponds to each possible combination of dice rolls.

## Pwgen
The `pwgen` tool located in cmd/pwgen provided in this project allows you to generate diceware passwords using the command line.

## Generating passwords
To generate a diceware password using the pwgen command-line tool, you can run the go run command followed by the path to the pwgen package and any options you want to use. By default, pwgen will generate a 6-word diceware password.

You can specify a different password length by using the -l or --length option followed by a number. For example, to generate a 10-word diceware password, you can run the following command:

```bash
go run github.com/8ff/diceware/pwgen -l 10
```

The resulting password will be printed to the console.

![pwgen demo](media/pwgen.gif)