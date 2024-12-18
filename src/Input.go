package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input takes a prompt and returns user input
func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
