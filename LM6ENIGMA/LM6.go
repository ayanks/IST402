package main

import (
	"fmt"
)

// Rotor represents an individual rotor
type Rotor struct {
	wiring   [26]int
	position int
	notch    int
}

// RotorSet represents a set of rotors in the Enigma machine
type RotorSet struct {
	rotors []*Rotor
}

// Plugboard represents the plugboard configuration
type Plugboard struct {
	mapping map[int]int
}

// Reflector represents the fixed reflector
type Reflector struct {
	wiring [26]int
}

// EnigmaMachine represents the entire Enigma machine setup
type EnigmaMachine struct {
	rotorSet  *RotorSet
	plugboard *Plugboard
	reflector *Reflector
}

// InitializeEnigmaMachine initializes an Enigma machine with provided settings
func InitializeEnigmaMachine() *EnigmaMachine {
	return &EnigmaMachine{
		rotorSet: &RotorSet{
			rotors: []*Rotor{
				{wiring: [26]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 1, 2, 3}, position: 0, notch: 1},
				{wiring: [26]int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 1, 2, 3, 4, 5, 6, 7}, position: 0, notch: 1},
			},
		},
		plugboard: &Plugboard{
			mapping: map[int]int{1: 2, 3: 4, 5: 6},
		},
		reflector: &Reflector{
			wiring: [26]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
		},
	}
}

// Encrypt encrypts a single character using the Enigma machine
func (em *EnigmaMachine) Encrypt(char rune) rune {
	intChar := int(char - 'A')

	// Advance rotor positions before encryption
	for _, rotor := range em.rotorSet.rotors {
		rotor.position = (rotor.position + 1) % 26
	}

	// Forward pass through plugboard mapping
	if mappedChar, exists := em.plugboard.mapping[intChar]; exists {
		intChar = mappedChar
	}

	// Forward pass through rotors
	for _, rotor := range em.rotorSet.rotors {
		intChar = (intChar + rotor.position) % 26
		intChar = rotor.wiring[intChar]
		intChar = (intChar - rotor.position + 26) % 26
	}

	// Pass through reflector
	intChar = em.reflector.wiring[intChar]

	// Backward pass through rotors
	for i := len(em.rotorSet.rotors) - 1; i >= 0; i-- {
		for j, value := range em.rotorSet.rotors[i].wiring {
			if value == (intChar+em.rotorSet.rotors[i].position)%26 {
				intChar = j
				break
			}
		}
		intChar = (intChar - em.rotorSet.rotors[i].position + 26) % 26
	}

	// Backward pass through plugboard mapping
	for mappedChar, originalChar := range em.plugboard.mapping {
		if intChar == mappedChar {
			intChar = originalChar
			break
		}
	}

	// Convert integer back to character representation
	encryptedChar := rune(intChar + 'A')
	return encryptedChar
}

// Decrypt decrypts a single character using the Enigma machine
func (em *EnigmaMachine) Decrypt(char rune) rune {
	intChar := int(char - 'A')

	// Backward pass through plugboard mapping
	for originalChar, mappedChar := range em.plugboard.mapping {
		if intChar == mappedChar {
			intChar = originalChar
			break
		}
	}

	// Backward pass through rotors
	for i := len(em.rotorSet.rotors) - 1; i >= 0; i-- {
		for j, value := range em.rotorSet.rotors[i].wiring {
			if j == intChar {
				intChar = (value + em.rotorSet.rotors[i].position) % 26
				break
			}
		}
		intChar = (intChar - em.rotorSet.rotors[i].position + 26) % 26
	}

	// Pass through reflector
	intChar = em.reflector.wiring[intChar]

	// Forward pass through rotors
	for _, rotor := range em.rotorSet.rotors {
		intChar = (intChar + rotor.position) % 26
		for j, value := range rotor.wiring {
			if j == intChar {
				intChar = (value - rotor.position + 26) % 26
				break
			}
		}
	}

	// Forward pass through plugboard mapping
	if originalChar, exists := em.plugboard.mapping[intChar]; exists {
		intChar = originalChar
	}

	// Convert integer back to character representation
	decryptedChar := rune(intChar + 'A')
	return decryptedChar
}

func main() {
	// Initialize Enigma machine
	enigma := InitializeEnigmaMachine()

	// Test encryption and decryption
	plaintext := "ISTTEAMSIX"
	ciphertext := ""
	decryptedText := ""

	// Encrypt
	for _, char := range plaintext {
		ciphertext += string(enigma.Encrypt(char))
	}

	// Decrypt
	for _, char := range ciphertext {
		decryptedText += string(enigma.Decrypt(char))
	}

	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Decrypted text:", decryptedText)
}
