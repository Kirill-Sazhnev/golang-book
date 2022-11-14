package main

import "fmt"

type celsius float64

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32)
}

type kelvin float64

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15) // Необходима конвертация типа
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var f fahrenheit = 227 // Аргумент должен быть типа kelvin
	k := f.kelvin()
	c := f.celsius()
	fmt.Print(c, "° C is ", k, "° K and ", f, " Fah") // Выводит: 294° K is 20.850000000000023° C
}

/*
Напишите программу с типами celsius, fahrenheit и kelvin и
методами для конвертации из одного типа температуры в другой.
*/
