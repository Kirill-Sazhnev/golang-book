package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {

	s = strings.ToLower(s)
	words := strings.Fields(s)
	var wordCount = make(map[string]int, len(words))

	for i, word := range words {
		words[i] = strings.Trim(word, ".,;—")
		wordCount[words[i]]++
	}
	return wordCount
}

func main() {

	var rawText = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear —except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

	countMap := wordCount(rawText)

	for word, count := range countMap {
		if count > 1 {
			fmt.Printf("Word \"%s\" repeats %d times\n", word, count)
		}
	}
}

/*
Напишите функцию для подсчета частоты упоминания слов в строке текста и
возвращения карты со словами и числом, указывающем, сколько раз они
употребляются. Функция должна конвертировать текст в нижний регистр и
обрезать знаки препинания. Используйте пакет strings. Функции, которые
пригодятся для выполнения данного задания: Fields, ToLower и Trim.

Используйте функцию для подсчета частоты слов следующего отрывка текста и
последующего вывода числа употребления каждого слова, что встречается более
одного раза.
*/
