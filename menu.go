package main

import "fmt"

func main() {

	balance := 0

	for {
		fmt.Println("\n--- МЕНЮ УПРАВЛЕНИЯ ---")
		fmt.Println("1. Зарабоок")
		fmt.Println("2. Проверить кошелек")
		fmt.Println("3. Рискованные инвестиции")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)
		// Программа замрет и будет ждать, пока ты нажмешь Enter

		switch choice {
		case 1:
			balance += 100
			fmt.Println(">>> Вы продали финики! Баланс пополнен на 100 ")
		case 2:
			fmt.Printf(">>> В вашем кошельке %d руб.\n", balance)
		case 3:
			if balance >= 200 {
				balance -= 200
				fmt.Println("Инвестиция в новые саженцы прошла успешно")
			} else {
				fmt.Println("Недостаточно средств для инвестиции!")
			}
		case 0:
			fmt.Println("Завершение работы... До свидания!")
			return // Это слово полностью остановит функцию main и выйдет из программы
		default:
			fmt.Println("Ошибка: вы ввели неверную команду!")
		}
	}
}
