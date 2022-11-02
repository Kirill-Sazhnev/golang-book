package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var coin float64
	var storage float64

	for storage < 20.0 {

		switch rand.Intn(3) {
		case 0:
			coin = 0.05
		case 1:
			coin = 0.10
		case 2:
			coin = 0.25
		}
		storage += coin
		fmt.Printf("Current balance is $%05.2f\n", storage)
	}

}

/*Напишите программу, которая случайным образом размещает монеты
пять ($0.05), десять ($0.10) и двадцать пять ($0.25) центов в пустую
копилку до тех пор, пока внутри не будет хотя бы двадцать
долларов ($20.00). Пускай после каждого пополнения копилки текущий
баланс отображается на экране,
отформатированный с нужной шириной и точностью.
*/
