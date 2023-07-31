package main

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/chacha20"
)

func encrypt(key, nonce, plaintext []byte) ([]byte, error) {
	//  key is exactly 32 bytes long
	if len(key) != chacha20.KeySize {
		return nil, fmt.Errorf("encryption error: invalid key size")
	}

	// Ensure that the nonce is exactly 12 bytes long
	if len(nonce) != chacha20.NonceSize {
		return nil, fmt.Errorf("encryption error: invalid nonce size")
	}

	// Creating a  ChaCha20 cipher with the provided key n nonce
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	// store the encrypted data
	encrypted := make([]byte, len(plaintext))

	// Encrypt the plaintext
	cipher.XORKeyStream(encrypted, plaintext)

	return encrypted, nil
}

func generateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func main() {
	// Generate a random key
	key, err := generateRandomBytes(chacha20.KeySize)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Generate a random nonce with 12 bytes
	nonce, err := generateRandomBytes(chacha20.NonceSize)
	if err != nil {
		fmt.Println("Error generating nonce:", err)
		return
	}

	// Get user input for plaintext
	var plaintext string
	fmt.Print("Enter plaintext to encrypt: ")
	fmt.Scanln(&plaintext)

	// Encrypt the plaintext
	encrypted, err := encrypt(key, nonce, []byte(plaintext))
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	fmt.Printf("Encrypted: %x\n", encrypted)

	// Decrypt the ciphertext
	decrypted, err := encrypt(key, nonce, encrypted)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Decrypted:", string(decrypted))
}
