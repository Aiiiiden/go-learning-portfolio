package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	balance, skill, gardenHealth, totalSales := loadGame()
	rand.Seed(time.Now().UnixNano())
	showGreetings()
	marketPrice := 100

	for {
		fmt.Println("\n--- МЕНЮ УПРАВЛЕНИЯ ---")
		marketPrice = 100 + (rand.Intn(41) - 20)
		fmt.Printf("ТЕКУЩАЯ РЫНОЧНАЯ ЦЕНА: %d руб. за порцию\n", marketPrice)
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
			balance, gardenHealth = sellDates(balance, skill, gardenHealth, marketPrice)
			balance = payTaxes(balance, gardenHealth)
			totalSales++
		case 2:
			printBalance(balance)
			fmt.Printf("Всего совершено продаж: %d\n", totalSales)
		case 3:
			balance, skill = invest(balance, skill)
		case 4:
			calculateIncome()
		case 5:
			balance, gardenHealth = fertilizeGarden(balance, gardenHealth)
		case 0:
			saveGame(balance, skill, gardenHealth, totalSales)
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

func loadGame() (int, int, int, int) {
	// Читаем файл
	data, err := os.ReadFile("save.txt")

	// Если файла нет (первый запуск), возвращаем стандартные значения
	if err != nil {
		return 0, 1, 100, 0
	}

	var b, s, h, t int
	// Вытаскиваем числа из прочитанных данных
	fmt.Sscanf(string(data), "%d %d %d %d", &b, &s, &h, &t)

	fmt.Println(">>> Прогресс успешно загружен!")
	return b, s, h, t
}

func sellDates(balance int, skill int, health int, price int) (int, int) {
	profit := price * skill
	event := rand.Intn(10)

	if event == 0 {
		fmt.Println("БЛАГОДАТЬ: Сугодняотличная погода, финики проданы дороже (х2)")
		profit = profit * 2
	} else if event == 1 {
		fmt.Println("ИСПЫТАНИЯ: Налетели вредители,часть урожая потеряна (х0.5)")
		profit = profit / 2
	}

	if health == 0 {
		fmt.Println("КОТОСТРОФА: Сад полностью засох вы ничего не собрали!")
		profit = 0
	} else if health < 50 {
		fmt.Println("!!! Сад истощен, урожай скудный.")
		profit = profit / 2
	}

	newhealth := health - 10
	if newhealth < 0 {
		newhealth = 0
	}
	fmt.Printf(">>> Вы продали финики (Уровень %d, Здоровье сада %d%%)! Цена за ед.: %d руб. Доход %d руб.\n", skill, health, price, profit)
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

func saveGame(balance int, skill int, health int, total int) {
	//создаем строку и наших данных
	data := fmt.Sprintf("%d %d %d %d", balance, skill, health, total)
	//записываем ее в текстовы файл
	//0644 - это права доступа стандарт для файлов
	err := os.WriteFile("save.txt", []byte(data), 0644)

	if err != nil {
		fmt.Println("Ошибка при схранении:", err)
	} else {
		fmt.Println(">>> Прогресс сохранен в файл!")
	}
}
