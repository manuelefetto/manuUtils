package ManuUtils

import (
	"fmt"
)

type Select struct {
	elements []string
	pointer  int
}

func GenerateSelect(elem []string) *Select {

	p := new(Select)
	(*p).elements = elem
	return p
}

func (sel *Select) renderElements() {

	for i, v := range sel.elements {
		if i == sel.pointer {
			fmt.Print("\x1b[36m") //set color to cyan
			fmt.Println(fmt.Sprintf("> %.33s", v))
			fmt.Print("\x1b[0m")
		} else {
			fmt.Println(fmt.Sprintf("  %.33s", v))
		}

	}
}

func (sel Select) deleteRenderedElements() {
	for i := 0; i < len(sel.elements); i++ {
		ClearLine()
	}
}

func (sel *Select) getInput() bool { // return true if user enter enter

	for {

		k := GetKeyboardKey()

		if k == 65517 {
			if sel.pointer == 0 {
				(*sel).pointer = len(sel.elements) - 1
			} else {
				(*sel).pointer = sel.pointer - 1
			}
			return false
		}

		if k == 65516 {
			if sel.pointer == len(sel.elements)-1 {
				(*sel).pointer = 0
			} else {
				(*sel).pointer = sel.pointer + 1
			}
			return false

		}

		if k == 13 {
			return true
		}
	}

}

func (sel Select) GetSelectedElement() string {

	return sel.elements[sel.pointer]

}

func (sel *Select) Select() string {

	var cicle bool = true

	for cicle {
		sel.renderElements()
		cicle = !sel.getInput()
		sel.deleteRenderedElements()
	}


	return sel.GetSelectedElement()

}
