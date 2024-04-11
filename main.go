package main

import (
	"fmt"
	Utils "manuUtils/packages"
	"time"
)

type Test struct {
	id      int
	nome    string
	cognome string
	cap     int
	prezzo  float64
	prova   struct {
		x string
		y int
	}
}

type Test2 struct {
	ezio, greggio, negro string
}

func main() {

	s2 := make([]Test2, 0)
	s2 = append(s2, Test2{"aaa", "aaa", "aaa"})
	s2 = append(s2, Test2{"bbb", "bbb", "bbb"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})
	s2 = append(s2, Test2{"ccc", "ccc", "ccc"})

	Utils.GenerateTable(s2, 5)

	// Loading()

	// s := make([]Test, 0)
	// s = append(s, Test{1, "aaa", "bbb", 28066, 23.23, struct {
	// 	x string
	// 	y int
	// }{"Akadddadadanandad", 2}})

	// Utils.GenerateTable(s)

	var stop string
	fmt.Scanln(&stop)
}

func Loading() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, "%")
		time.Sleep(time.Millisecond * 300)
		Utils.ClearLine()
	}
}

func testSelect() {
	selectElement := [...]string{"Manuele", "Lara", "Davide"}

	s := Utils.GenerateSelect(selectElement[:])

	scelta := s.Select()

	fmt.Println("Scelto :", scelta)
}

func testCheck() {

	s := Utils.GenerateCheckbox([]string{"Lara", "Larax", "Larastella", "Manu", "Manux", "Manuelefetto"}, false)

	fmt.Println(s.Select())
}
