package main

import (
	"fmt"
	"testing"
)

func TestSortType(t *testing.T) {
	deck := CardDeckConstructor().sort()
	if CheckType(deck) != "*main.CardDeck" {
		t.Log("Метод sort возвращает неверный тип")
		t.Fail()
	}
}
func TestConstructType(t *testing.T) {
	deck := CardDeckConstructor()
	if CheckType(deck) != "*main.CardDeck" {
		t.Log("Метод CardDeckConstructor возвращает неверный тип")
		t.Fail()
	}

}

func CheckType(v *CardDeck) string {
	return fmt.Sprintf("%T", v)
}
