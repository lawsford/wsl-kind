package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 'y' to proceed or 'n' to cancel: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}
	inputTrimmed := strings.TrimSpace(strings.ToLower(input))
	return inputTrimmed == "y"
}