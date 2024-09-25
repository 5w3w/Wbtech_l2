package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определяем флаги
	fieldsFlag := flag.String("f", "", "Выбрать поля (колонки), например: 1,2,3")
	delimiterFlag := flag.String("d", "\t", "Разделитель для полей (по умолчанию TAB)")
	separatedFlag := flag.Bool("s", false, "Только строки с разделителем")

	// Разбираем флаги
	flag.Parse()

	if *fieldsFlag == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: необходимо указать -f для выбора колонок.")
		os.Exit(1)
	}

	// Парсим запрошенные колонки
	fieldIndices, err := parseFields(*fieldsFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка в формате полей:", err)
		os.Exit(1)
	}

	// Сканируем строки с stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверяем, содержит ли строка разделитель, если установлен флаг -s
		if *separatedFlag && !strings.Contains(line, *delimiterFlag) {
			continue
		}

		// Разбиваем строку по разделителю
		columns := strings.Split(line, *delimiterFlag)

		// Собираем и выводим запрошенные поля
		output := make([]string, 0, len(fieldIndices))
		for _, index := range fieldIndices {
			if index > 0 && index <= len(columns) {
				output = append(output, columns[index-1])
			}
		}

		if len(output) > 0 {
			fmt.Println(strings.Join(output, *delimiterFlag))
		}
	}

	// Проверка на ошибки при сканировании
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
		os.Exit(1)
	}
}

// parseFields парсит строку полей в список индексов.
func parseFields(fields string) ([]int, error) {
	var result []int
	parts := strings.Split(fields, ",")
	for _, part := range parts {
		var index int
		_, err := fmt.Sscanf(part, "%d", &index)
		if err != nil {
			return nil, fmt.Errorf("некорректный формат поля: %s", part)
		}
		result = append(result, index)
	}
	return result, nil
}
