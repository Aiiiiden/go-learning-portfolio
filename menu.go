package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	balance, skill, gardenHealth, totalSales, medjulCount, regularCount, wildCount, medjulMoney := loadGame()
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
		fmt.Println("6. Просмотр статистики")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)
		// Программа замрет и будет ждать, пока ты нажмешь Enter

		switch choice {
		case 1:
			balance, gardenHealth, medjulCount, regularCount, wildCount, medjulMoney = sellDates(balance, skill, gardenHealth, marketPrice, medjulCount, regularCount, wildCount, medjulMoney)
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
		case 6:
			fmt.Println("\n========== СТАТИСТИКА САДА ==========")
            fmt.Printf("Королевский Меджул:  %d сборов (Выручка: %d руб.)\n", medjulCount, medjulMoney)
            fmt.Printf("Обычные финики:      %d сборов\n", regularCount)
            fmt.Printf("Дикие финики (лечебные): %d сборов\n", wildCount)
            fmt.Println("=====================================")

			if medjulCount> regularCount{
				fmt.Println("Ваша стратегия: Рискованный Султан!")
			}else if wildCount> regularCount {
				fmt.Println("Ваша стратегия: Мудрый Садовник!")
			} else {
				fmt.Println("Ваша стратегия: Стабильный Фермер.")
			}
			
		case 0:
			saveGame(balance, skill, gardenHealth, totalSales, medjulCount, regularCount, wildCount, medjulMoney)
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

func loadGame() (int, int, int, int, int, int, int, int) {
	// Читаем файл
	data, err := os.ReadFile("save.txt")

	// Если файла нет (первый запуск), возвращаем стандартные значения
	if err != nil {
		return 0, 1, 100, 0,0,0,0,0
	}

	var b, s, h, t, mc, rc, wc, mm int
	// Вытаскиваем числа из прочитанных данных
	fmt.Sscanf(string(data), "%d %d %d %d %d %d %d %d", &b, &s, &h, &t, &mc, &rc, &wc, &mm)

	fmt.Println(">>> Прогресс успешно загружен!")
	return b, s, h, t, mc, rc, wc, mm
}

func sellDates(balance, skill, health, price, medjulC, regC, wildC,medjulM int) (int, int, int, int, int, int) {
	fmt.Println("\n--- ВЫБОР СОРТА ДЛЯ СБОРА ---")
	fmt.Println("1. Обычные (Стандартный дохо, износ -10%)")
	fmt.Println("2. Меджул (Доход х3, износ -35%!)")
	fmt.Println("3. Дикие (Доход почти 0, ЛЕЧИТ САД +5%)")
	fmt.Print("ВАШ ВЫБОР:")

	var cropChoice int
	fmt.Scan(&cropChoice)
	
	profit := price * skill
	wear := 10

	switch cropChoice {
	case 1:
		fmt.Println(">>>Сщбираем обычные финики...")
		regC++
	case 2:
		fmt.Println(">>>РИСК! Собираем элитный Меджул...")
		profit = profit * 3
		wear = 35
		medjulC++
		medjulM += profit
	case 3:
		fmt.Println(">>> Забота о саде: собираем дикие финики...")
		profit = 10
		wear = -5
		wildC++
	default:
		fmt.Println("Неверный выбор собираем как обычно.")
	}

	event := rand.Intn(10)

	if event == 0 {
		fmt.Println("БЛАГОДАТЬ: Сегодня отличная погода, финики проданы дороже (х2)")
		profit *= 2
	} else if event == 1 {
		fmt.Println("ИСПЫТАНИЯ: Налетели вредители,часть урожая потеряна (х0.5)")
		profit /= 2
	}

	if health <= 0 {
		fmt.Println("КОТОСТРОФА: Сад полностью засох вы ничего не собрали!")
		profit = 0
	} else if health < 50 && cropChoice != 3 {
		fmt.Println("!!! Сад истощен, урожай скудный.")
		profit /= 2
	}

	newHealth := health - wear
	if newHealth > 100 {newHealth = 100}
	if newHealth < 0   {newHealth = 0}

	fmt.Printf(">>> Итог: Доход %d руб. Здоровье сада теперь: %d%%\n", profit, newHealth)
	return balance + profit, newHealth, medjulC, regC, wildC, medjulM
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

func saveGame(balance, skill, health, total, medjulC, regC, wildC, medjulM int) {
	//создаем строку и наших данных
	data := fmt.Sprintf("%d %d %d %d %d %d %d %d", balance, skill, health, total, medjulC, regC, wildC, medjulM)
	//записываем ее в текстовы файл
	//0644 - это права доступа стандарт для файлов
	err := os.WriteFile("save.txt", []byte(data), 0644)

	if err != nil {
		fmt.Println("Ошибка при схранении:", err)
	} else {
		fmt.Println(">>> Прогресс сохранен в файл!")
	}
}
