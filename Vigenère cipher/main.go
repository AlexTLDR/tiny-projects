// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher

package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	fmt.Println(decipher(cipherText, keyword))

	message := "your message goes here"
	keyword = "golang"
	fmt.Println(cipher(message, keyword))
}

func decipher(cipherText string, keyword string) string {
	var message string
	var keyIndex int

	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// cipher letter - key letter
		c = (c-k+26)%26 + 'A'
		message += string(c)

		// increment keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}

	return message
}

func cipher(message string, keyword string) string {
	var cipherText string
	var keyIndex int

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := 0; i < len(message); i++ {
		// A=0, B=1, ... Z=25
		m := message[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// message letter + key letter
		c := (m+k)%26 + 'A'
		cipherText += string(c)

		// increment keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}

	return cipherText
}
