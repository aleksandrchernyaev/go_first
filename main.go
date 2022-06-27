package main

import (
	"fmt"
)

func main() {

	message := "Скоро я стану ниндзя"

	fmt.Println(message)
	ChangeMessege(&message)

	//Первый комментарий
	fmt.Println(message)

	//Второй комментарий
	fmt.Println(message)
	fmt.Println(message2)
}

func ChangeMessege(message *string) {

	*message += " (из функции ChangeMessege)"
}
