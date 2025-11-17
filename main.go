package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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

	fmt.Printf("Ваш выбор %s\n", ELEMENTS[playerSelect])

	fmt.Print("Сейчас компьютер выберет свой элемент")
	for key := 0; key < 5; key++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\n")

	computerSelect := rand.Intn(2) + 1
	fmt.Printf("Компьютер выбрал %s\n", ELEMENTS[computerSelect])
	fmt.Print("\n")
	fmt.Printf("")
}

// func getWinner(playerSelect int, computerSelect int) {
// 	switch 
// }