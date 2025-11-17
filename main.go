package main

import (
	"fmt"
	"os"
)

func main() {
	ELEMENTS := map[int]string{
		1: "Камень",
		2: "Ножницы",
		3: "Бумага",
	}

	fmt.Println("Добро пожаловать в игру Камень-ножницы-бумага. Автор Lorriant. Помощник Артур")

	fmt.Println("Выберите ваш элемент. Нужно ввести номер.")
	for key := 1; key < 4; key++ {
		fmt.Printf("%d. %s\n", key, ELEMENTS[key])
	}

	var playerSelect int
	for {
		count, err := fmt.Scanf("%d", &playerSelect)
		if (err != nil || count != 1 || playerSelect < 1 || playerSelect > 3) {
			fmt.Fprint(os.Stderr, "Ты тупой? Нужно ввести число 1-3\n")
			continue
		} else {
			break
		}
	}

	fmt.Printf("Ваш выбор №: %s\n", ELEMENTS[playerSelect])
}
