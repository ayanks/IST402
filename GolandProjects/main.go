// CFB Cipher Feedback Mode
package main

import (
	"fmt"
)

/* an array with 4 rows and 2 columns*/
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}
var cipher = [4]int{}
var iv int = 0b10

func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}
func main() {
	var xor int = 0
	var lookupValue int = 0
	lookupValue = codebookLookup(iv)
	//Display the original Message
	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value of a is %02b\n", message[i])
	}
	//Ciphertext
	for i := 0; i < 4; i++ {
		xor = message[i] ^ lookupValue
		lookupValue = codebookLookup(xor)
		fmt.Printf("The ciphered value of a is %02b\n", xor)
		cipher[i] += xor
	}
	//Plaintext
	lookupValue = codebookLookup(iv)
	for i := 0; i < 4; i++ {
		xor = cipher[i] ^ lookupValue
		lookupValue = codebookLookup(cipher[i])
		fmt.Printf("The plaintext value of a is %02b\n", xor)
	}
}
