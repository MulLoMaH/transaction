/*файл 1 (main.go) - главное меню:
-функция main (c циклом и рассылкой по другим функциям в зависимости от результата)
-функция mainMenu (хранит в себе пункты меню, принимает выбор и передает результат в main)*/

package main

import "fmt"

type Transaction struct {
	ID       int     `json:"ID"`                   // Уникальный номер (1, 2, 3...)`json:"ID операции"`
	Type     string  `json:"Поступление/Списание"` // "доход" или "расход"`json:"Поступление/Списание"`
	Amount   float64 `json:"Сумма"`                // Сумма (1500.50)`json:"Сумма операции"`
	Category string  `json:"Категория"`            // "оплата клиента", "налоги", "оборудование", "обучение"
	Note     string  `json:"Описание"`             // Описание ("За лендинг для Ивана")
	Date     string  `json:"Дата"`                 // Дата в формате "2026-01-30"
}

type Storage struct {
	transactions []Transaction
}

var transactions []Transaction
var nextID = 1

func main() {
	loadTransaction()
	for {
		choice := mainMenu()
		switch choice {
		case 1:
			newTransaction()
		case 2:
			allTransaction()
		case 3:
			balance()
		case 4:
			filterCategory()
		case 5:
			saveTransaction()
		case 6:
			exitTransaction()
		default:
			fmt.Println("Введено не корректное значение")

		}

	}

}

func mainMenu() int {
	fmt.Println("=== ФИНАНСОВЫЙ ТРЕКЕР ===")
	fmt.Println("1. Добавить транзакцию")
	fmt.Println("2. Показать все транзакции")
	fmt.Println("3. Показать баланс (доходы - расходы)")
	fmt.Println("4. Фильтр по категории")
	fmt.Println("5. Сохранить")
	fmt.Println("6. Выход")
	fmt.Println("Ваш выбор:")

	var choice int
	fmt.Scan(&choice)
	return choice

}
