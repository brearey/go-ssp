package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	ELEMENTS := map[int]string{
		1: "Камень",
		2: "Ножницы",
		3: "Бумага",
	}
	clearScreen()
	fmt.Println("Добро пожаловать в игру Камень-ножницы-бумага. Автор Lorriant. Помощник Артур")

	fmt.Println("Выберите ваш элемент. Нужно ввести номер.")
	for key := 1; key < 4; key++ {
		fmt.Printf("%d. %s\n", key, ELEMENTS[key])
	}

	var playerSelect, computerSelect, count int
	var s string
	var scanErr, convErr error
	for {
		count, scanErr = fmt.Scan(&s)
		playerSelect, convErr = strconv.Atoi(s)
		if (scanErr != nil || convErr != nil || count != 1 || playerSelect < 1 || playerSelect > 3) {
			clearScreen()
			fmt.Fprint(os.Stderr, "Ты тупой? Нужно ввести число 1-3\n")
			continue
		} else {
			break
		}
	}
	clearScreen()
	fmt.Printf("Ваш выбор %s\n", ELEMENTS[playerSelect])

	fmt.Print("Сейчас компьютер выберет свой элемент")
	for key := 0; key < 5; key++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\n")

	computerSelect = rand.Intn(2) + 1
	fmt.Printf("Компьютер выбрал %s\n", ELEMENTS[computerSelect])
	fmt.Print("\n")
	fmt.Printf("")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// func getWinner(playerSelect int, computerSelect int) {
// 	switch 
// }