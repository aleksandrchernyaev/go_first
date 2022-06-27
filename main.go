package main

import (
	"fmt"
)

func main() {

	message := "Скоро я стану ниндзя"

	fmt.Println(message)
	ChangeMessege(&message)
	fmt.Println(message)
}

func ChangeMessege(message *string) {

	*message += " (из функции ChangeMessege)"
}
