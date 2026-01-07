package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(11)
	var userGuess int

	attempts := 3

	fmt.Println("🎮 Добро пожаловать в игру 'Угадай число'!")

	fmt.Printf("У вас есть %d попытки чтобы угадать число от 0 до 10\n", attempts)

	for attempts > 0 {
		fmt.Printf("\nПопыток осталось: %d. Ваше число: ", attempts)

		_, err := fmt.Scanln(&userGuess)
		if err != nil {
			fmt.Println("Ошибка чтения ввода. Пожалуйста, введите одно целое число от 0 до 10.")
			continue
		}

		if userGuess < 0 || userGuess > 10 {
			fmt.Println("❌ Ошибка! Число должно быть от 0 до 10")
			continue
		}

		if userGuess == secretNumber {
			fmt.Println("🎉 Поздравляем! Вы угадали!")
			break
		} else if userGuess > secretNumber {
			fmt.Printf("Неверно! Загаданное число меньше чем %d\n", userGuess)
		} else {
			fmt.Printf("Неверно! Загаданное число больше чем %d\n", userGuess)
		}

		attempts--
	}

	if attempts <= 0 {
		fmt.Printf("\n💔 Вы проиграли! Загаданное число было: %d\n", secretNumber)
	}
}
