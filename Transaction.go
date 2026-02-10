/*
файл 1 (main.go) - главное меню:
-функция main (c циклом и рассылкой по другим функциям в зависимости от результата)
-функция mainMenu (хранит в себе пункты меню, принимает выбор и передает результат в main)

файл 2 (addNew.go) - добавить транзакцию: (попробовать реализовать через горутины и каналы)
-функция newTransaction()
-функция addType
-функция addAmount
-функция addCategory
-функция addNote
-функция addDate

файл 3 (actions.go) - показать все транзакции и баланс:
-функция allTransaction()
-функция balanse()


файл 4 (filter.go) - фильтр по категории:
-функция filterCategory()
-функция ...


файл 5 (saveExit.go) - сохранить и выйти:
-функция saveTransaction()
-функция saveJSON()
-функция saveCSV()
-функция loadTransaction()
-функция exitTransaction()



//main.go

package main

type Transaction struct {
	ID        int `json:"ID"`    // Уникальный номер (1, 2, 3...)`json:"ID операции"`
    Type      string `json:"Поступление/Списание"` // "доход" или "расход"`json:"Поступление/Списание"`
    Amount    float64 `json:"Сумма"`// Сумма (1500.50)`json:"Сумма операции"`
    Category  string `json:"Категория"` // "оплата клиента", "налоги", "оборудование", "обучение"
    Note      string `json:"Описание"` // Описание ("За лендинг для Ивана")
    Date      string `json:"Дата"`// Дата в формате "2026-01-30"
}

type Storage struct {
	transactions[]Transaction
}

var transactions[]Transaction
var nextID = 1

func main() {
	loadTransaction()
	for {
		choice := mainMenu()
		switch choise {
		case 1:
			newTransaction()
		case 2:
			allTransaction()
		case 3:
			balanse()
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

func mainMenu()int {
	var choice int
	fmt.Scan(&choice)
	fmt.Println("=== ФИНАНСОВЫЙ ТРЕКЕР ===")
	fmt.Println("1. Добавить транзакцию")
	fmt.Println("2. Показать все транзакции")
	fmt.Println("3. Показать баланс (доходы - расходы)")
	fmt.Println("4. Фильтр по категории")
	fmt.Println("5. Сохранить")
	fmt.Println("6. Выход")
	fmt.Println("Ваш выбор:")
	return choice

}


























func main() {
	fmt.Print("Программа перезапустится через: ")
	for i := 3; i >= 0; i-- {
	fmt.Printf("\r%d", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("\r")

	}
}*/





 




