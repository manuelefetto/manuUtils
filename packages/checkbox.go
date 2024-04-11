package ManuUtils

import "fmt"

type Element struct {
	element string
	checked bool
}

type Checkbox struct {
	elements []Element
	pointer  int
}

func GenerateCheckbox(elements []string, allChecked bool) *Checkbox {

	checkbox := new(Checkbox)

	checkbox.elements = func() []Element {

		elementsElements := make([]Element, 0)

		for _, v := range elements {
			var tmp Element
			tmp.element = v
			tmp.checked = allChecked
			elementsElements = append(elementsElements, tmp)
		}

		return elementsElements

	}()

	checkbox.pointer = 0

	return checkbox

}

func (checkbox *Checkbox) renderCheckbox() {

	for i, v := range checkbox.elements {
		if i == checkbox.pointer {
			fmt.Print("\x1b[36m") //set color to cyan
			fmt.Println(fmt.Sprintf("> [%s]%.33s", func() string {
				if v.checked {
					return "x"
				} else {
					return " "
				}
			}(), v.element))
			fmt.Print("\x1b[0m")
		} else {
			fmt.Println(fmt.Sprintf("  [%s]%.33s", func() string {
				if v.checked {
					return "x"
				} else {
					return " "
				}
			}(), v.element))
		}
	}

}

func (checkbox Checkbox) deleteRenderedElements() {
	for i := 0; i < len(checkbox.elements); i++ {
		ClearLine()
	}
}

func (checkbox *Checkbox) getInput() bool { // return true if user enter enter

	for {

		k := GetKeyboardKey()

		if k == 65517 {
			if checkbox.pointer == 0 {
				(*checkbox).pointer = len(checkbox.elements) - 1
			} else {
				(*checkbox).pointer = checkbox.pointer - 1
			}
			return false
		}

		if k == 65516 {
			if checkbox.pointer == len(checkbox.elements)-1 {
				(*checkbox).pointer = 0
			} else {
				(*checkbox).pointer = checkbox.pointer + 1
			}
			return false

		}

		if k == 32 {
			(*checkbox).elements[checkbox.pointer].checked = !checkbox.elements[checkbox.pointer].checked
			return false
		}

		if k == 13 {
			return true
		}
	}

}

func (checkbox Checkbox) GetSelectedElement() []string {

	selected := make([]string, 0)

	for _, v := range checkbox.elements {
		if v.checked {
			selected = append(selected, v.element)
		}
	}

	return selected
}

func (checkbox *Checkbox) Select() []string {

	var cicle bool = true

	for cicle {
		checkbox.renderCheckbox()
		cicle = !checkbox.getInput()
		checkbox.deleteRenderedElements()
	}

	return checkbox.GetSelectedElement()

}
