package main

import "testing"

func TestWrongInputPosition(t *testing.T) {
	pos := " h44 t"
	moves := checkMovements(pos)
	s := ""
	for _, v := range moves {
		s += v
	}
	if s != "g6f5g2f3" {
		t.Log("Набор ходов неверен!")
		t.Fail()
	}
}

func TestMovements(t *testing.T) {
	arrayoftests := [4]int{0, 1, 2, 6}
	for i, v := range arrayoftests {
		if checkMargin(v) != v%(i+1) {
			t.Log("Ошибка расчета отступов от края доски!")
			t.Fail()
		}
	}
}
