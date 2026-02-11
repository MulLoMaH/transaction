/*файл 4 (filter.go) - фильтр по категории:
-функция filterCategory()
-функция ...*/

package main

import (
	"fmt"
	"strings"
)

func filterCategory() {
	var category string
	fmt.Print("Введите категорию для фильтрации: ")
	fmt.Scan(&category)

	//Счетчик найденных транзакций
	count := 0

	//Нормализуем регистр для сравнения
	normalizedinput := strings.ToLower(category)

	//Фильтрация и вывод транзакций
	for _, transaction := range transactions {
		normalizedCategory := strings.ToLower(transaction.Category)

		if normalizedCategory == normalizedinput {

			//Выводим категорию транзакции один раз (только для первой транзакции)
			if count == 0 {
				fmt.Printf("\n=== Категория: %s ===\n", transaction.Category)
			}

			//Выводим каждую транзакцию
			fmt.Printf("\nТранзакция #%d:\n", transaction.ID)
			fmt.Printf("Тип: %s\n", transaction.Type)
			fmt.Printf("Сумма: %.2f\n", transaction.Amount)
			fmt.Printf("Описание: %s\n", transaction.Note)
			fmt.Printf("Дата: %s\n", transaction.Date)

			count++
		}
	}

	//Проверка на пустой результат
	if count == 0 {
		fmt.Printf("\nТранзакции в категории '%s' не найдены\n", category)
	} else {
		fmt.Printf("\nНайдено %d транзакций в категории %s\n", count, category)
	}
}
