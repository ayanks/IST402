package main

import (
	"fmt"
	"strings"
)

// This is the EnigmaMachine represents the Enigma-like machine.
type EnigmaMachine struct {
	rotor      string
	plugboard  map[rune]rune
	reflectors map[rune]rune
}

// creates a new Enigma-like machine with specified rotor and plugboard settings.
func NewEnigmaMachine(rotor string, plugboardSettings map[rune]rune) *EnigmaMachine {
	// Simplified rotor configuration, where each character is shifted by 1.
	rotor = strings.ToUpper(rotor)

	// Default reflector configuration.
	reflectors := map[rune]rune{
		'A': 'Z', 'B': 'Y', 'C': 'X', 'D': 'W', 'E': 'V', 'F': 'U',
		'G': 'T', 'H': 'S', 'I': 'R', 'J': 'Q', 'K': 'P', 'L': 'O',
		'M': 'N', 'N': 'M', 'O': 'L', 'P': 'K', 'Q': 'J', 'R': 'I',
		'S': 'H', 'T': 'G', 'U': 'F', 'V': 'E', 'W': 'D', 'X': 'C',
		'Y': 'B', 'Z': 'A',
	}

	return &EnigmaMachine{
		rotor:      rotor,
		plugboard:  plugboardSettings,
		reflectors: reflectors,
	}
}

// EncryptChar encrypts a single character.
func (e *EnigmaMachine) EncryptChar(char rune) rune {
	// Apply plugboard substitution, if applicable.
	if val, ok := e.plugboard[char]; ok {
		char = val
	}

	// Passes through the rotor.
	rotorIndex := strings.IndexRune(e.rotor, char)
	if rotorIndex != -1 {
		char = rune((rotorIndex+1)%26 + 'A')
	}

	//substitutes reflector
	if val, ok := e.reflectors[char]; ok {
		char = val
	}

	// Reverse rotor operation
	rotorIndex = strings.IndexRune(e.rotor, char)
	if rotorIndex != -1 {
		char = rune((rotorIndex+25)%26 + 'A')
	}

	return char
}

// EncryptString encrypts string using the code
func (e *EnigmaMachine) EncryptString(input string) string {
	encrypted := ""
	for _, char := range strings.ToUpper(input) {
		if char >= 'A' && char <= 'Z' {
			encryptedChar := e.EncryptChar(char)
			encrypted += string(encryptedChar)
		}
	}
	return encrypted
}

func main() {
	// Create a new Enigma-like machine with rotor settings and plugboard
	rotorSettings := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	plugboardSettings := map[rune]rune{
		'A': 'B', 'C': 'D', 'E': 'F', 'G': 'H', 'I': 'J', 'K': 'L',
		'M': 'N', 'O': 'P', 'Q': 'R', 'S': 'T', 'U': 'V', 'W': 'X',
		'Y': 'Z',
	}
	enigma := NewEnigmaMachine(rotorSettings, plugboardSettings)

	// Input message to encrypt.
	message := "IST402TEAM6"

	// Encrypt the message.
	encryptedMessage := enigma.EncryptString(message)
	fmt.Printf("Original Message: %s\n", message)
	fmt.Printf("Encrypted Message: %s\n", encryptedMessage)
}
