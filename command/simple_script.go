package command

import (
	"bytes"
	"fmt"
	"regexp"
)

type simpleScript struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(simpleScript)
	commandList[command.GetSignature()] = command
}

func (simpleScript simpleScript) GetSignature() string {
	return "simpleScript"
}

func (simpleScript simpleScript) GetDescription() string {
	return "this is a Description"
}

func (simpleScript simpleScript) Handle() {
	tokens := tokenizer("add( 1 + \"100\" )")
	fmt.Println()
	fmt.Println(tokens)
}

type tokenItem struct {
	Type  string `json:"type" simple:"sType"`
	Value string `json:"value" simple:"sValue"`
}

func tokenizer(input string) []tokenItem {
	tokens := make([]tokenItem, 0)

	current := 0
	for current < len(input) {
		char := input[current]
		fmt.Printf("%q", char)
		if char == '(' {
			tokens = append(tokens, tokenItem{"paren", "("})
		}
		if char == ')' {
			tokens = append(tokens, tokenItem{"paren", ")"})
		}
		WHITESPACE := regexp.MustCompile(`\s`)
		if WHITESPACE.MatchString(string(char)) {
			current++
			continue
		}

		NUMBERS := regexp.MustCompile(`[0-9]`)
		if NUMBERS.MatchString(string(char)) {
			var numbersValue bytes.Buffer

			for NUMBERS.MatchString(string(char)) {
				numbersValue.WriteString(string(char))
				current++
				char = input[current]
			}
			tokens = append(tokens, tokenItem{"number", numbersValue.String()})
		}

		if string(char) == "\"" {
			var stringValue bytes.Buffer
			current++
			char = input[current]
			strChar := string(char)
			for strChar != "\"" {
				stringValue.WriteString(strChar)
				current++
				char = input[current]
				strChar = string(char)
			}

			tokens = append(tokens, tokenItem{"string", stringValue.String()})
			current++
			continue
		}

		LETTERS := regexp.MustCompile(`[a-z]`)
		if LETTERS.MatchString(string(char)) {
			var lettersValue bytes.Buffer

			for LETTERS.MatchString(string(char)) {
				lettersValue.WriteString(string(char))
				current++
				char = input[current]
			}
			tokens = append(tokens, tokenItem{"name", lettersValue.String()})
			continue
		}

		current++
	}

	return tokens
}

// func simpleParser(tokens []tokenItem) {
// 	current := 0
// 	walk := func() map[string]string {
// 		token := tokens[current]

// 		for token.Type != "paren" || (token.Type == "paren" && token.Value != ")") {

// 		}
// 		if token.Type == "number" {
// 			current++
// 			return map[string]string{
// 				"type":  "NumberLiteral",
// 				"value": token.Value,
// 			}
// 		}
// 		if token.Type == "string" {
// 			current++
// 			return map[string]string{
// 				"type":  "StringLiteral",
// 				"value": token.Value,
// 			}
// 		}
// 		// if token.Type == "paren" && token.Value == "("{
// 		// 	current++
// 		// 	token = tokens[current]

// 		// 	node := map[string]
// 		// }
// 		return map[string]string{}
// 	}

// 	return
// }
