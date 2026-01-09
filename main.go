package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(11)

	attempts := 3

	fmt.Println("🎮 Добро пожаловать в игру 'Угадай число'!")

	fmt.Printf("У вас есть %d попытки чтобы угадать число от 0 до 10\n", attempts)

	for attempts > 0 {
		fmt.Printf("\nПопыток осталось: %d. Ваше число: ", attempts)

		userGuess, err := readUserInput()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		if userGuess == secretNumber {
			fmt.Println("🎉 Поздравляем! Вы угадали!")
			break
		}

		giveHint(userGuess, secretNumber)

		attempts--
	}

	if attempts <= 0 {
		fmt.Printf("\n💔 Вы проиграли! Загаданное число было: %d\n", secretNumber)
	}
}

func readUserInput() (int, error) {
	var userGuess int
	_, err := fmt.Scanln(&userGuess)

	if err != nil {
		return 0, fmt.Errorf("нужно ввести целое число")
	}

	if userGuess < 0 || userGuess > 10 {
		return 0, fmt.Errorf("число должно быть от 0 до 10")
	}

	return userGuess, nil
}

func giveHint(guess, secret int) {
	if guess > secret {
		fmt.Printf("Неверно! Загаданное число меньше чем %d\n", guess)
	} else {
		fmt.Printf("Неверно! Загаданное число больше чем %d\n", guess)
	}
}
