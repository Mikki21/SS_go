/*
Техно фичи ГО
1. Мульти парадигменный : процедурная, функциональная, ОО.
2. Многопоток. Concurrency
3. Инструменты


Типы данных:
	Примитивные:
		bool true/false
		Числа:
			целые:
				int | int8, int16, int32, int64
				uint | uint8=byte
			Вещественные:
				float32
				float64
			Комплексные:
				complex64
				complex128
		Symbols:
			rune (16 ричное )(uint16) ex: 'a'
		Strings:
			string ex: `a` or "a"   ! `` doesnt read unicode symbols
				1) \a \b etc...
				2)\u \x \o
Placeholders in GOLNAG
	Uses with fmt.Printf("string",p), p is declared.
		%v - prints the value.
		%+v - if value is structured, shows name of lines in structure
		%#v - prints a Go syntax representation of the value
		%T - prints type of the value
		%t -
		%d - prints decimal value
		%b - prints binary value
		%c - prints char equivalent
		%x - 16x value
		%f - basic decimal formatting of float.
		%e %E - форматирует флоат в научной нотации
		%s - простой вывод строки
		%q
		%x (for string)
		%p - прдставление ссылки.
		%6.2f - ширина вывода и точность





*/
package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {

	//p := point{1, 2}
	fmt.Printf("%s\n", "\"string\"")
}
