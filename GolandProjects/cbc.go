package main

import (
	"fmt"
)

/* an array with 4 rows and 2 columns*/
var codebookcbc = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var message1 = [4]int{0b00, 0b01, 0b10, 0b11}
var ciphertxt = [4]int{}
var i int = 0b10

func codebookLookupCBC(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebookcbc[i][j] == xor {
			j++
			lookupValue = codebookcbc[i][j]
			break
		}
	}
	return lookupValue
}

func main() {
	var xor int = 0
	var lookupValue int = 0
	lookupValue = codebookLookupCBC(i)

	//Display the original Message
	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value of a is %02b\n", message1[i])
	}

	//Ciphertext
	for i := 0; i < 4; i++ {
		xor = message1[i] ^ lookupValue
		lookupValue = codebookLookupCBC(xor)
		fmt.Printf("The ciphered value of a is %02b\n", xor)
		ciphertxt[i] = xor
	}

	//Plaintext
	lookupValue = codebookLookupCBC(i)
	for i := 0; i < 4; i++ {
		xor = ciphertxt[i] ^ lookupValue
		lookupValue = codebookLookupCBC(ciphertxt[i])
		fmt.Printf("The plaintext value of a is %02b\n", xor)
	}
}
