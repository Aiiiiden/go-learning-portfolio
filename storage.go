package main

import "fmt"

type Palm struct {
	Type string
	Age  int
}

func main() {
	myStorage := make(map[string]Palm)

	myStorage["p1"] = Palm{Type: "Аджва", Age: 10}
	myStorage["p2"] = Palm{Type: "Меджул", Age: 5}
	myStorage["p3"] = Palm{Type: "Суккари", Age: 8}

	if value, ok := myStorage["p5"]; ok { // Специально ищем p5, которой нет
    fmt.Printf("Пальма найдена: %s, Возраст: %d\n", value.Type, value.Age)
	} else {
    fmt.Println("Ошибка: Пальма под номером p5 не числится в реестре!")
	}

	if value, ok := myStorage["p2"]; ok { // Специально ищем p2, которой нет
    fmt.Printf("Пальма найдена: %s, Возраст: %d\n", value.Type, value.Age)
	} else {
    	fmt.Println("Ошибка: Пальма под номером p2 не числится в реестре!")
	}
}