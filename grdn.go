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

func (p *Palm) Water() {
	p.Health += 10
	fmt.Printf("Вы полили сорт %s. Здоровье теперь %d%%\n", p.Type, p.Health)
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
			myGarden[i].Water()
		}
	}

	//Итогоый отчёт после всех изменений
	fmt.Println("\n--- ИТОГОВЫЙ ОТЧЁТ ЧЕРЕЗ 3 ГОДА ---")
	PrintGardenReport(myGarden)

	fmt.Println("\nБисмиллягь симляция завершена.")
}

func PrintGardenReport(garden []Palm) {
	totalYield := 0
	for _, p := range garden {
		current := p.CalculateActualYield()
		fmt.Printf("Сорт: %-8s | Возраст: %d лет | Здоровье: %d%% | Урожа: %d кг\n",
			p.Type, p.Age, p.Health, current)
		totalYield += current
	}

	fmt.Printf("ОБЩИЙ УРОЖАЙ САДА: %d кг\n", totalYield)
}
