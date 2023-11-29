// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher

package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	for {
		for i := 0; i < len(cipherText); i++ {
			c := cipherText[i]
			if c >= 'a' && c <= 'z' {
				c -= 3
				if c < 'a' {
					c += 26
				}
			} else if c >= 'A' && c <= 'Z' {
				c -= 3
				if c < 'A' {
					c += 26
				}
			}
			fmt.Printf("%c", c)
		}
	}
}
