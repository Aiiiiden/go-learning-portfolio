package main

import "fmt"

func main() {

	balance := 0
	showGreetings()

	for {
		fmt.Println("\n--- МЕНЮ УПРАВЛЕНИЯ ---")
		fmt.Println("1. Зарабоок")
		fmt.Println("2. Проверить кошелек")
		fmt.Println("3. Рискованные инвестиции")
		fmt.Println("4. Рассчитать рогноз дохода")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)
		// Программа замрет и будет ждать, пока ты нажмешь Enter

		switch choice {
		case 1:
			balance = sellDates(balance)
		case 2:
			printBalance(balance)
		case 3:
			balance = invest(balance)
		case 4:
			calculateIncome()
		case 0:
			fmt.Println("Завершение работы... До свидания!")
			return // Это слово полностью остановит функцию main и выйдет из программы
		default:
			fmt.Println("Ошибка: вы ввели неверную команду!")
		}
	}
}

func showGreetings() {
	fmt.Println("Добро пожаловать в систеу управления садом, Амир!")
}

func sellDates(balance int) int {
	balance += 100
	fmt.Println(">>> Вы продали финики! Баланс пополнен на 100 ")
	return balance
}

func printBalance(currentBalance int) {
	fmt.Printf("--- ТЕКУЩИЙ СЧЁТ: %d руб. ---\n", currentBalance)
}

func invest(balance int) int {
	if balance >= 200 {
		fmt.Println("Инвестиция в новые саженцы прошла успешно")
		return balance - 200
	} else {
		fmt.Println("Недостаточно средств для инвестиции!")
		return balance
	}
}

func calculateIncome() {
	var palmCount int
	var harvestPerPalm int
	var price int

	fmt.Println("Сколько всего пальм в саду? ")
	fmt.Scan(&palmCount)
	fmt.Println("Сколько кг фиников даёт одна пальма? ")
	fmt.Scan(&harvestPerPalm)
	fmt.Println("Цена за 1 кг фиников: ")
	fmt.Scan(&price)

	totalProfit := palmCount * harvestPerPalm * price
	fmt.Printf("Твой прогноз дохода %d руб.\n", totalProfit)
}
