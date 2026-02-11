/*файл 3 (actions.go) - показать все транзакции и баланс:
-функция allTransaction()
-функция balanсe()*/

package main

import (
	"fmt"
	"strings"
)

func allTransaction() {
	for i := range transactions {
		fmt.Println(transactions[i])
	}
}

func balance() {
	sum := 0.0
	for i := range transactions {
		// Tagged switch: сравниваем значение сразу в заголовке
		switch strings.ToLower(transactions[i].Type) {
		case "поступление":
			sum += transactions[i].Amount
		case "списание":
			sum -= transactions[i].Amount
			// default: ничего не делаем (игнорируем неизвестные типы)
		}
	}
	fmt.Printf("Баланс: %.2f\n", sum)
}
