package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Тип для сортировки строк по месяцу
var monthOrder = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
	"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

func main() {
	// Определение флагов
	colFlag := flag.Int("k", 1, "Указание колонки для сортировки")
	numFlag := flag.Bool("n", false, "Сортировать по числовому значению")
	reverseFlag := flag.Bool("r", false, "Сортировать в обратном порядке")
	uniqueFlag := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	monthFlag := flag.Bool("M", false, "Сортировать по названию месяца")
	ignoreTrailingSpacesFlag := flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	checkSortedFlag := flag.Bool("c", false, "Проверить отсортированы ли данные")
	humanNumericFlag := flag.Bool("h", false, "Сортировать по числовому значению с суффиксами")

	flag.Parse()

	// Открытие входного файла
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Чтение строк
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Удаление хвостовых пробелов, если установлен флаг -b
	if *ignoreTrailingSpacesFlag {
		for i := range lines {
			lines[i] = strings.TrimRight(lines[i], " ")
		}
	}

	// Сортировка по ключам
	if *uniqueFlag {
		lines = unique(lines)
	}

	if *checkSortedFlag {
		if isSorted(lines, *numFlag, *monthFlag, *humanNumericFlag, *colFlag) {
			fmt.Println("Файл отсортирован.")
		} else {
			fmt.Println("Файл не отсортирован.")
		}
		return
	}

	sort.SliceStable(lines, func(i, j int) bool {
		// Сортировка по колонке -k
		vi := getColumn(lines[i], *colFlag)
		vj := getColumn(lines[j], *colFlag)

		// Если установлен флаг -M, сортируем по месяцу
		if *monthFlag {
			mi, mj := monthOrder[vi], monthOrder[vj]
			if *reverseFlag {
				return mj < mi
			}
			return mi < mj
		}

		// Если установлен флаг -n, числовая сортировка
		if *numFlag {
			ni, _ := strconv.Atoi(vi)
			nj, _ := strconv.Atoi(vj)
			if *reverseFlag {
				return nj < ni
			}
			return ni < nj
		}

		// Если установлен флаг -h, сортировка с суффиксами
		if *humanNumericFlag {
			return compareHumanNumeric(vi, vj, *reverseFlag)
		}

		// Стандартная строковая сортировка
		if *reverseFlag {
			return vj < vi
		}
		return vi < vj
	})

	// Запись результата в файл или вывод
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Функция для извлечения нужной колонки
func getColumn(line string, col int) string {
	parts := strings.Fields(line)
	if col <= len(parts) {
		return parts[col-1]
	}
	return ""
}

// Удаление дубликатов
func unique(lines []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if _, value := keys[line]; !value {
			keys[line] = true
			result = append(result, line)
		}
	}
	return result
}

// Проверка отсортированности
func isSorted(lines []string, numFlag, monthFlag, humanNumericFlag bool, colFlag int) bool {
	for i := 1; i < len(lines); i++ {
		if !compare(lines[i-1], lines[i], numFlag, monthFlag, humanNumericFlag, colFlag, false) {
			return false
		}
	}
	return true
}

// compare - сравнение двух строк с учетом флагов
func compare(a, b string, numFlag, monthFlag, humanNumericFlag bool, colFlag int, reverse bool) bool {
	// Получаем нужные колонки
	colA := getColumn(a, colFlag)
	colB := getColumn(b, colFlag)

	// Если установлен флаг -M, сравниваем по месяцу
	if monthFlag {
		monthA, monthB := monthOrder[colA], monthOrder[colB]
		if reverse {
			return monthA > monthB
		}
		return monthA < monthB
	}

	// Если установлен флаг -n, сравниваем числовые значения
	if numFlag {
		numA, errA := strconv.Atoi(colA)
		numB, errB := strconv.Atoi(colB)
		if errA != nil || errB != nil {
			return false
		}
		if reverse {
			return numA > numB
		}
		return numA < numB
	}

	// Если установлен флаг -h, используем сравнение с суффиксами
	if humanNumericFlag {
		return compareHumanNumeric(colA, colB, reverse)
	}

	// Стандартное строковое сравнение
	if reverse {
		return colA > colB
	}
	return colA < colB
}

// compareHumanNumeric - сравнение чисел с суффиксами
func compareHumanNumeric(a, b string, reverse bool) bool {
	// Преобразуем строки с суффиксами в числа для сравнения
	valA, multiplierA := parseHumanNumeric(a)
	valB, multiplierB := parseHumanNumeric(b)

	// Умножаем значения на их множители (для "К", "М", "Г" и т.д.)
	valA *= multiplierA
	valB *= multiplierB

	// Сравниваем числа
	if reverse {
		return valA > valB
	}
	return valA < valB
}

// parseHumanNumeric - парсит строку с суффиксами и возвращает числовое значение и множитель
func parseHumanNumeric(s string) (float64, float64) {
	// Определение множителя
	multipliers := map[string]float64{
		"K": 1e3,  // килобайты (тысячи)
		"M": 1e6,  // мегабайты (миллионы)
		"G": 1e9,  // гигабайты (миллиарды)
		"T": 1e12, // терабайты (триллионы)
		"P": 1e15, // петабайты
		"E": 1e18, // эксабайты
	}

	// Определение числового значения и суффикса
	var numberPart string
	var suffixPart string
	for i, ch := range s {
		if ch >= '0' && ch <= '9' || ch == '.' {
			numberPart += string(ch)
		} else {
			suffixPart = s[i:]
			break
		}
	}

	// Парсим числовую часть
	value, err := strconv.ParseFloat(numberPart, 64)
	if err != nil {
		return 0, 1 // Возвращаем 0 в случае ошибки
	}

	// Получаем множитель
	if multiplier, ok := multipliers[suffixPart]; ok {
		return value, multiplier
	}

	// Если нет суффикса, возвращаем 1 как множитель
	return value, 1
}
