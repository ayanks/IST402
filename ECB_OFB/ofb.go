package main

import (
	"fmt"
)

/* an array with 4 rows and 2 columns*/
var c = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var messageofb = [4]int{0b00, 0b01, 0b10, 0b11}
var cipher1 = [4]int{}
var v int = 0b10

func codebookLookupOFB(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if c[i][j] == xor {
			j++
			lookupValue = c[i][j]
			break
		}
	}
	return lookupValue
}

func main() {
	var xor int
	var lookupValue int

	// Initialize OFB
	lookupValue = codebookLookupOFB(v)

	// Display the original Message
	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value of a is %02b\n", messageofb[i])
	}

	// Encryption (Ciphertext)
	for i := 0; i < 4; i++ {
		// Encrypt using the lookup value (feedback)
		xor = messageofb[i] ^ lookupValue
		fmt.Printf("The ciphered value of a is %02b\n", xor)
		cipher1[i] = xor

		// Update the lookupValue for the next round
		lookupValue = codebookLookupOFB(lookupValue)
	}

	// Decryption (Plaintext)
	// Reset lookupValue to the initial value (iv)
	lookupValue = codebookLookupOFB(v)
	for i := 0; i < 4; i++ {
	}
}
