package main

import (
	"fmt"
	"strings"
)

func main() {
	//cipherText := "ECFRZKYGLGRMUSDHRXK"
	plainText := "your message goes here"
	keyword := "GOLANG"

	for i, letter := range sypher(plainText, keyword) {
		i %= len(keyword)

		Desyph := letter - rune(keyword[i])
		Desyph = (Desyph + 26) % 26

		fmt.Printf("%c", Desyph+'A')
	}

}

func sypher(t, k string) string {
	message := ""
	text := strings.Replace(t, " ", "", -1)
	text = strings.ToUpper(text)

	for i, letter := range text {
		i %= len(k)

		Desyph := letter + rune(k[i])
		Desyph = (Desyph) % 26

		message += fmt.Sprintf("%c", Desyph+'A')
	}
	return message
}

/*
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// зашифрованная буква - ключевая буква
		c = (c-k+26)%26 + 'A'
		message += string(c)

		// увеличение keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}

	fmt.Println(message)
*/
