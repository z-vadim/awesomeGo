package main

import "fmt"

func main() {
	mWord := "radar"
	caesarCipher(mWord)
	//if isPalindrome(mWord) != true {
	//	fmt.Printf("Word '%v' - isn't palindrome", mWord)
	//} else {
	//	fmt.Printf("Word '%v' - is palindrome", mWord)
	//}
}

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encrypts a plaintext string by shifting each character with the provided key.
func EncryptPlaintext(plaintext string, key int) string {
	return rotateText(plaintext, key)
}

// Decrypts a ciphertext string by reverse shifting each character with the provided key.
func DecryptCiphertext(ciphertext string, key int) string {
	return rotateText(ciphertext, -key)
}

// Takes a string and rotates each character by the provided amount.
func caesarCipher(str string) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+rot)%26]
		}
	}
	return string(rotatedText)
}
