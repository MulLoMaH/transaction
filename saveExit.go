package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func saveTransaction() {
	fmt.Println("В какой формат сохранить файл?")
	fmt.Println("1.JSON")
	fmt.Println("2.CSV")
	fmt.Println("3.Сохранить в оба формата")
	fmt.Println("4.Выход без/после сохранения")

	fmt.Print("Введите цифру: ")

	var choice int

	fmt.Scan(&choice)
	switch choice {
	case 1:
		saveJSON(transactions)
	case 2:
		saveCSV(transactions)
	case 3:
		saveJSON(transactions)
		saveCSV(transactions)
	case 4:
		exitTransaction()
	default:

	}
}

func saveJSON(transactions []Transaction) {
	file, err := os.Create("transaction.json")
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", " ")

	encoder.Encode(transactions)
	fmt.Println("✅ Сохранено в transaction.json")
}

func saveCSV(transactions []Transaction) {
	file, err := os.Create("transaction.csv")
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}

	defer file.Close()

	file.Write([]byte{0xEF, 0xBB, 0xBF})

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	defer writer.Flush()

	writer.Write([]string{
		"ID",
		"Поступление/Списание",
		"Сумма",
		"Категория",
		"Описание",
		"Дата"})
	for _, t := range transactions {
		writer.Write([]string{
			strconv.Itoa(t.ID),
			t.Type,
			strconv.FormatFloat(t.Amount, 'f', 2, 64),
			t.Category,
			t.Note,
			t.Date,
		})
		fmt.Println("✅ Сохранено в transaction.csv")
	}

}

func loadTransaction() {
	data, err := os.ReadFile("transaction.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("ℹ️ Файл с транзакциями не найден — создаём новый")
			transactions = []Transaction{}
			nextID = 1
		} else {
			fmt.Println("❌ Ошибка чтения файла:", err)
		}
		return
	}

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		fmt.Println("❌ Ошибка загрузки данных:", err)
		return
	}
	// Восстанавливаем следующий ID
	if len(transactions) > 0 {
		nextID = transactions[len(transactions)-1].ID + 1
	}
	fmt.Printf("✅ Загружено %d транзакций\n", len(transactions))
}

func exitTransaction() {
	fmt.Print("Выход из программы")

	// Обратный отсчёт: 3 → 2 → 1 секунды
	for seconds := 3; seconds > 0; seconds-- {
		// Анимация точек внутри каждой секунды: . → .. → ...
		for dots := 1; dots <= 3; dots++ {
			fmt.Printf("\rВыход через %d сек%-3s", seconds, strings.Repeat(".", dots))
			time.Sleep(300 * time.Millisecond)
		}
	}

	fmt.Println("\rВыход через 0 сек... 👋 До свидания!")
	os.Exit(0)
}
