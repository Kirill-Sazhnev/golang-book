package main

import (
	"fmt"
	"math/rand"
)

type temperature float32

func fahToCel(f temperature) temperature {
	return temperature((f - 32.0) * 5.0 / 9.0)
}

func celToFah(c temperature) temperature {
	return temperature(c*9.0/5.0 + 32)
}

type tempConv func(f temperature) temperature

func drawTable(c tempConv, from, to string) tempConv {
	fmt.Println("=======================")
	fmt.Printf("| °%-7v | °%-7v |\n", from, to)
	fmt.Println("=======================")

	return func(f temperature) temperature {
		return c(f)
	}
}

func main() {
	converter := drawTable(fahToCel, "F", "C")
	for i := 0; i < 7; i++ {
		temp := temperature(rand.Intn(30) * 5)
		fmt.Printf("| %-8.2v | %-8.2f |\n", temp, converter(temp))
	}
	fmt.Println("=======================")

	converter = drawTable(celToFah, "C", "F")
	for i := 0; i < 7; i++ {
		temperature := temperature(rand.Intn(13)+8) * 5
		fmt.Printf("| %-8.2v | %-8.2f |\n", temperature, converter(temperature))
	}
	fmt.Println("=======================")
}

/*
Программа должна построить две таблицы. В первой таблице два столбца, в
первом значится температура по Цельсию °C, а во втором — по Фаренгейту °F.
Значения должны быть от 40° C до 100° C шагами в 5°. Для заполнения столбцов
 требуется использовать методы конвертации, описанные в уроке о методах.

После заполнения одной таблицы заполните вторую таким образом, чтобы столбцы
были инвертированы. То есть конвертация должна проводиться из градусов по
Фаренгейту в градусы по Цельсию.

Код, что вы напишите для создания таблицы, в будущем можно будет использовать
вновь, уже для других программ, содержимое которых нужно отобразить в таблице
с двумя столбцами. Используйте функции для разделения кода который создает
таблицы от кода для вычисления значений температуры каждой строки.

Реализуйте функцию drawTable, что принимает функцию первого класса в качестве
параметра и вызывает ее для получения данных каждой созданной строки.
Результатом передачи другой функции к drawTable должны быть другие
отображаемые данные.
*/
