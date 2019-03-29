package util

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

type condition func(string) bool

func PromptForInput(f condition, text string) string {
	input := ""

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
		return PromptForInput(f, text)
	} else if (f != nil && f(string(sentence))) {
		fmt.Println(text)
		return PromptForInput(f, text)
	}

	input = string(sentence)
	input = strings.TrimSpace(input)
	return input
}
