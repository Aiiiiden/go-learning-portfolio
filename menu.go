package main

import (
	"fmt"
)

func main() {

	balance := 0
	skill := 1
	gardenHealth := 100
	showGreetings()

	for {
		fmt.Println("\n--- МЕНЮ УПРАВЛЕНИЯ ---")
		fmt.Println("1. Зарабоок")
		fmt.Println("2. Проверить кошелек")
		fmt.Println("3. Рискованные инвестиции")
		fmt.Println("4. Рассчитать рогноз дохода")
		fmt.Println("5. Уход за садом")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)
		// Программа замрет и будет ждать, пока ты нажмешь Enter

		switch choice {
		case 1:
			balance, gardenHealth = sellDates(balance, skill, gardenHealth)
			balance = payTaxes(balance, gardenHealth)
		case 2:
			printBalance(balance)
		case 3:
			balance, skill = invest(balance, skill)
		case 4:
			calculateIncome()
		case 5:
			balance, gardenHealth = fertilizeGarden(balance, gardenHealth)
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

func sellDates(balance int, skill int, health int) (int, int) {
	profit := 100 * skill
	if health < 50 {
		fmt.Println("!!! Сад истощен, урожай скудный.")
		profit = profit / 2
	}

	newhealth := health - 10
	if newhealth < 0 {
		newhealth = 0
	}
	fmt.Printf(">>> Вы продали финики (Уровень %d, Здоровье сада %d%%)! Доход %d руб.\n", skill, health, profit)
	return balance + profit, newhealth
}

func printBalance(currentBalance int) {
	fmt.Printf("--- ТЕКУЩИЙ СЧЁТ: %d руб. ---\n", currentBalance)
}

func invest(balance int, skill int) (int, int) {
	if balance >= 300 {
		fmt.Println(">>>Вы прошли курсы садовода! Уровень повышен!")
		return balance - 300, skill + 1
	} else {
		fmt.Println("Недостаточно средств для обучения! (нужно 300 руб)")
		return balance, skill
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

func payTaxes(balance int, health int) int {
    if balance >= 1000 { 
        if health <= 20 {
            fmt.Println("Послабление: Закят не взымается из-за плохого состояния сада.")
            return balance
        }
        balance = balance - (balance / 10)
        fmt.Println("Вы выплатили Закят/Налог на развитие сада.")
    }
    return balance
}

func fertilizeGarden(balance int, health int) (int, int) {
	if balance >= 150 {
		fmt.Println("Вы купили лучшее удобрение! Сад снова цветёт.")
		return balance - 150, 100
	} else {
		fmt.Println(">>>Недостаточно средств для ухода за садом!")
		return balance, health
	}
}
