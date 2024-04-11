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

type Test3 struct {
	nome           string
	cognome        string
	numeroTelefono string
	pene           float64
}

func main() {

	slice := []Test3{
		{"Mario", "Rossi", "123456789", 15.5},
		{"Luigi", "Verdi", "987654321", 20.3},
		{"Giovanni", "Bianchi", "456789123", 18.2},
		{"Paolo", "Neri", "789123456", 14.8},
		{"Antonio", "Gialli", "321654987", 17.6},
		{"Giuseppe", "Marroni", "654987321", 19.1},
		{"Marco", "Arancioni", "987123654", 16.3},
		{"Alessandro", "Verdi", "456321789", 22.7},
		{"Stefano", "Blu", "789654123", 13.4},
		{"Davide", "Rosa", "321789654", 21.0},
		{"Roberto", "Gialli", "654123789", 18.9},
		{"Simone", "Arancioni", "123654987", 16.8},
		{"Andrea", "Bianchi", "987321654", 14.2},
		{"Fabio", "Neri", "654789321", 23.5},
		{"Massimo", "Rossi", "321987654", 12.7},
		{"Federico", "Blu", "789654321", 20.6},
		{"Gianluca", "Marroni", "456123789", 17.9},
		{"Emanuele", "Verdi", "987321456", 15.3},
		{"Claudio", "Rosa", "654789123", 19.8},
		{"Nicola", "Bianchi", "321456987", 18.4},
	}

	Utils.GenerateTable(slice, 5)

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
