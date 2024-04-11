package ManuUtils

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func ClearCLI() {
	fmt.Print("\033[2J\033[H")
}

func ClearLine() {
	fmt.Print("\033[2K\r\033[F\033[2K\r")
}

func GetKeyboardKey() int {
	if err := keyboard.Open(); err != nil {
		panic(err) //TODO migliorare la gestione degli errori
	}
	_, k, _ := keyboard.GetKey()

	return int(k)
}
