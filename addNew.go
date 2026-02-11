/*файл 2 (addNew.go) - добавить транзакцию: (попробовать реализовать через горутины и каналы)
-функция newTransaction()
-функция addType
-функция addAmount
-функция addCategory
-функция addNote
-функция addDate*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func newTransaction() {
	newType := addType()
	newAmount := addAmount()
	newCategory := addCategory()
	newNote := addNote()
	newDate := addDate()

	transaction := Transaction{
		ID:       nextID,
		Type:     newType,
		Amount:   newAmount,
		Category: newCategory,
		Note:     newNote,
		Date:     newDate,
	}

	transactions = append(transactions, transaction)
	nextID++

	fmt.Printf("\nТранзакция добавлена:\nТип: %s\nСумма: %.2f\nКатегория: %s\nОписание: %s\nДата: %s\n",
		newType, newAmount, newCategory, newNote, newDate)

}

func addType() string {
	var newType string
	for {
		fmt.Print("Укажите, было Поступление или Списание: ")
		fmt.Scan(&newType)
		if newType == "Поступление" || newType == "поступление" || newType == "Списание" || newType == "списание" {
			return newType
		} else {
			fmt.Println("Введено не коректное значение: ", newType)
		}
	}
}

func addAmount() float64 {
	var newAmount float64
	for {
		fmt.Print("Укажите сумму транзакции: ")
		_, err := fmt.Scan(&newAmount)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			fmt.Scanln() // очистить буфер ввода
			continue
		}
		if newAmount <= 0 {
			fmt.Println("Сумма должна быть положительной")
		} else {
			return newAmount
		}
	}
}

func addCategory() string {
	var newCategory string
	for {
		fmt.Print("Укажите категорию транзакции: ")
		fmt.Scan(&newCategory)
		if newCategory == "" {
			fmt.Println("Категория не указана")
		} else {
			return newCategory
		}
	}
}

func addNote() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Укажите описание транзакции: ")
		newNote, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		newNote = strings.TrimSpace(newNote) // убрать \n и пробелы по краям
		if newNote == "" {
			fmt.Println("Описание не указано")
		} else {
			return newNote
		}
	}
}

func addDate() string {
	return time.Now().Format("02.01.2006")
}
