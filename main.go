package main

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

func main() {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader("Johan"))
	if err != nil {
		return
	}
	fmt.Printf("Morse Code is: %s\n", string(morseCode))
}
