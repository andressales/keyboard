// Package keyboard reads user input from the keyboard
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboard
// It returns the number read and any error encountered
func GetFloat() (input float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	textValue, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	textValue = strings.TrimSpace(textValue)
	input, err = strconv.ParseFloat(textValue, 64)
	if err != nil {
		return 0, err
	}

	return input, nil
}
