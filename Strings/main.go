package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {

	message := "Hola Estación Espacial Internacional"

	for _, letter := range message {
		if letter >= 'a' && letter <= 'z' {
			letter += 13
			if letter < 'a' {
				letter -= 26
			}
		} else if letter >= 'A' && letter <= 'Z' {
			letter += 13
			if letter < 'A' {
				letter -= 26
			}
		}
		fmt.Printf("%c", letter)
	}

}

/*Расшифруйте цитату Юлия Цезаря: L fdph, L vdz, L frqtxhuhg.

Ваша программа должна будет сдвинуть буквы верхнего и нижнего
регистра на -3. Помните, что ‘a’ становится ‘x’, ‘b’ становится
‘y’, а ‘c’ становится ‘z’. То же самое происходит с буквами
верхнего регистра.

Зашифруйте сообщение на испанском: “Hola Estación Espacial
Internacional”  через ROT13. Модифицируйте Листинг 7 с
использованием ключевого слова range. Теперь, когда вы
используете ROT13 c испанским текстом, ударение над буквами
сохраняется.


*/
