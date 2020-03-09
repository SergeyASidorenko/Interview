// Copyright 2020 The Home. All rights not reserved.
// Пакет с реализацией тестового задание на сортировку карточной колоды
// и печатью в консоль результата
// Сведения о лицензии отсутствуют
package main

import "fmt"

// CardDeck Структура карточной колоды. Это отображение порядкового номера карты в колоде на структуру, содержащую
// два элемента, каждый из которых это массив. Первый элемент это масть->значение, второй - номинал->значение
// Колода сокращенная - 36 карт
type CardDeck struct {
	Deck   map[int]map[string]string
	Values map[string]int
	Suits  map[string]int
}

/*
Функция сортировки
Сортирует в порядке убывания, то есть номиналы выстраиваются так: от тузов к шестеркам, масти от червей - к пикам
Алгоритм таков. Пробегаем массив, делая перестановку пары карт, если текущая карта больше предыдущей, большая карта "всплывает" к началу колоды", проверка идет сразу по двум критериям
по масти и по номиналу,так как масть больший вес, то проверку по номиналу мы проверяем по условию, если текущая карта имеет больший номинал, чем предыдущая,
но перестанавливаем их мы только если у них одинаковые масти, чтобы не портить то, что было сделать перестановкой, выполненной предыдущим условием

*/
func (d *CardDeck) sort() *CardDeck {
	for g := 1; g < 36; g++ {
		for i := 1; i < 36; i++ {
			if d.Suits[d.Deck[i]["suit"]] > d.Suits[d.Deck[i-1]["suit"]] { // Проверка по масти
				tmp := d.Deck[i-1]
				d.Deck[i-1] = d.Deck[i]
				d.Deck[i] = tmp
			} else if d.Suits[d.Deck[i]["suit"]] == d.Suits[d.Deck[i-1]["suit"]] && d.Values[d.Deck[i]["value"]] > d.Values[d.Deck[i-1]["value"]] { // Проверка по номиналу, при условии, что масти одинаковы
				tmp := d.Deck[i-1]
				d.Deck[i-1] = d.Deck[i]
				d.Deck[i] = tmp
			}
		}
	}
	return d
}

//Печать колоды
func (d *CardDeck) print() {
	for i := 0; i < 36; i++ {
		fmt.Printf("Номинал %v, Значение %v \n", d.Deck[i]["suit"], d.Deck[i]["value"])
	}
}

//CardDeckConstructor Своего рода конструктор, заполнене колоды можно выполнить вызовом этой функции
func CardDeckConstructor() *CardDeck {
	obj := CardDeck{}
	obj.Deck = make(map[int]map[string]string, 36)
	obj.Values = map[string]int{"Six": 1, "Seven": 2, "Eight": 3, "Nine": 4, "Ten": 5, "Knave": 6, "Queen": 7, "King": 8, "Ace": 9}
	obj.Suits = map[string]int{"Shades": 1, "Crosses": 2, "Diamonds": 3, "Hearts": 4}

	//В этом цикле идет заполнение массива в порядке возрастания
	for i := 0; i < 36; i++ {
		innermap := make(map[string]string, 2)
		//int(i/9)+1 - выражение будет возвращать только числа в интервале 0-8,
		//оставаясь постоянным в течение 9 итераций, так как мы хотим заполнить колоду картами
		// в таком порядке сначала все номиналы одной масти, за тем все номиналы другой итд
		innermap["suit"] = getKey(obj.Suits, int(i/9)+1)
		//i%9+1 - выражение будет возвращать только числа в интервале 0-8,причем возврат от 8 к нулю будет осуществляться после очередных 9 итераций, так как мы их
		//отображаем на номиналы, используя CardDeck.Values
		innermap["value"] = getKey(obj.Values, i%9+1)
		obj.Deck[i] = innermap
	}
	return &obj
}

// Вспомогательная функция Поиск ключа по значению
// Вообще лучше ее сделать методом структуры CardDeck, так как она используется для нужд этого класса, но я пока не читал про это
func getKey(arr map[string]int, val int) string {
	for key, value := range arr {
		if val == value {
			return key
		}
	}
	return ""
}

func main() {

	d := CardDeckConstructor()
	//d.sort().print()
	d.print()

}
