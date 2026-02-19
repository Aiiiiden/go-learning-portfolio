package main

import (
	"fmt"
)

type Palm struct {
	Type   string
	Yield  int
	Health int
	Age    int
}

// Считаем урожай на снове возраста и здоровья
func (p Palm) CalculateActualYield() int {
	if p.Age < 5 {
		return 0 //Совсем маленкая не блодов
	}

	actualYield := p.Yield

	//Если дерево старое (старше 15 лет), урожай падает на 20%
	if p.Age > 15 {
		actualYield = int(float64(actualYield) * 0.8)
	}

	//Если здоровье ниже 70%, урожай падает ещё в двое
	if p.Health < 70 {
		actualYield = actualYield / 2
	}
	return actualYield
}

func (p *Palm) Water(balans *int) {
	if *balans >= 50 {
		*balans -= 50 
		p.Health += 10
		fmt.Printf("Вы полили %s за 50 руб. Здоровье: %d%%, Остаток:%d руб\n", p.Type, p.Health, *balans)
	} else {
		fmt.Printf("!!!Не хватает денег для полива %s (нужно 50, у вас %d)\n", p.Type, *balans)
	}
	

}

// PassYear симулирует прохождение одного года в саду
func PassYear(garden []Palm) {
	for i := range garden {
		garden[i].Age++       //	дерево стареет на год
		garden[i].Health -= 5 //    здоровье немного падает без ухода

		//Ограничитель. здоровье не может быть меньше 0
		if garden[i].Health < 0 {
			garden[i].Health = 0
		}
	}
}

func main() {

	balans := 0
	pricePerKg := 10
	
	fmt.Println("--- Добро пожаловать в твой Сад в Медине ---")
	
	//создаем начальный сад
	myGarden := []Palm{
		{Type: "Меджул", Yield: 50, Health: 80, Age: 6},
		{Type: "Аджва", Yield: 60, Health: 95, Age: 8},
		{Type: "Суккари", Yield: 40, Health: 70, Age: 7},
	}

	//сад до симуляции
	fmt.Println("\n--- СОСТОЯНИЕ САДА СЕЙЧАС ---")
	PrintGardenReport(myGarden)

	//Симуляция машина времени
	fmt.Println("\n... Прошло 3 года засухи и времен...\n")
	for year := 1; year <= 3; year++ {
		PassYear(myGarden) //каждый год дерево стареет и здоровье падает
	}

	//Проверка здоровья
	fmt.Println("--- ОБХОД САДА ЧЕРЕЗ 3 ЛЕТ ---")
	for i := range myGarden {
		if myGarden[i].Health < 50 {
			fmt.Printf("!!! Дерево %s ослабло (Health: %d%%). Срочно поливаем!\n",
				myGarden[i].Type, myGarden[i].Health)
			myGarden[i].Water(&balans)
		}
	}

	//Итогоый отчёт после всех изменений
	fmt.Println("\n--- ИТОГОВЫЙ ОТЧЁТ ЧЕРЕЗ 3 ГОДА ---")
	currentTotal := PrintGardenReport(myGarden)

	profit := currentTotal * pricePerKg
	fmt.Printf("Мы продали урожай и получили %d руб\n", profit)
	balans = balans + profit
	fmt.Printf("У нас в кошельке %d руб", balans)
	fmt.Println("\nБисмиллягь симляция завершена.")
}

func PrintGardenReport(garden []Palm) int {
	totalYield := 0
	for _, p := range garden {
		current := p.CalculateActualYield()
		fmt.Printf("Сорт: %-8s | Возраст: %d лет | Здоровье: %d%% | Урожа: %d кг\n",
			p.Type, p.Age, p.Health, current)
		totalYield += current
		}

	fmt.Printf("ОБЩИЙ УРОЖАЙ САДА: %d кг\n", totalYield)
	return  totalYield
}
