// Copyright 2020 The Home. All rights not reserved.
// Пакет с реализацией тестового задание на нахождение
// всевозможных ходов шахматного коня
// Пользователь вводит текущее положение фигуры через параметр
// Сведения о лицензии отсутствуют

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	checkMovements("h4")

}
func checkMovements(position string) []string {
	mappos := make(map[string]int, 8)
	margins := make(map[string]int, 4)
	/*
		Отображение вертикалей доски на числа, удобно выполнять арифметику после этого над вертикалями
	*/
	var movements []string
	mappos["a"] = 1
	mappos["b"] = 2
	mappos["c"] = 3
	mappos["d"] = 4
	mappos["e"] = 5
	mappos["f"] = 6
	mappos["g"] = 7
	mappos["h"] = 8
	/*
		простейшая обработка ошибок,
		строка очищается от начальных и конечных пробелов,
		в расчет используются толька два символа
	*/
	position = strings.Trim(position, " ")
	position = position[:2]

	var vertical string = string(position[0])   //вертикаль текущего положения коня
	var horizontal string = string(position[1]) //горизонталь текущего положения коня
	var ihorizontal int
	ihorizontal, _ = strconv.Atoi(horizontal)
	//находим отступы от краёв доски по всем сторонам
	//причем если отступ больше 2 клеток, все равно принимаем это расстояние как 2 клетки!
	margins["top"] = checkMargin(8 - ihorizontal)
	margins["left"] = checkMargin(mappos[vertical] - 1)
	margins["right"] = checkMargin(8 - mappos[vertical])
	margins["bottom"] = checkMargin(ihorizontal - 1)
	//Поиск всех возможных ходов

	for i := margins["top"]; i >= 1; i-- {
		if i*margins["left"] >= 2 {
			movements = append(movements, getKey(mappos, mappos[vertical]-(3-i))+strconv.Itoa(ihorizontal+i))
		}
		if i*margins["right"] >= 2 {
			movements = append(movements, getKey(mappos, mappos[vertical]+(3-i))+strconv.Itoa(ihorizontal+i))

		}
	}
	for i := margins["bottom"]; i >= 1; i-- {
		if i*margins["left"] >= 2 {
			movements = append(movements, getKey(mappos, mappos[vertical]-(3-i))+strconv.Itoa(ihorizontal-i))
		}
		if i*margins["right"] >= 2 {
			movements = append(movements, getKey(mappos, mappos[vertical]+(3-i))+strconv.Itoa(ihorizontal-i))

		}
	}
	fmt.Println(movements)
	return movements
}

// Вспомогательная функция расчета отступов
func checkMargin(value int) int {
	if value == 0 || value < 0 {
		return 0
	} else if value == 1 {
		return value
	} else {
		return 2
	}
}

// Вспомогательная функция получения ключа отображения по значению
func getKey(arr map[string]int, val int) string {
	for key, value := range arr {
		if val == value {
			return key
		}
	}
	return ""
}
