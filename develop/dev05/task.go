package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	after      = flag.Int("A", 0, "Print +N lines after the match")
	before     = flag.Int("B", 0, "Print +N lines before the match")
	context    = flag.Int("C", 0, "Print ±N lines around the match")
	countOnly  = flag.Bool("c", false, "Print count of matching lines")
	ignoreCase = flag.Bool("i", false, "Ignore case")
	invert     = flag.Bool("v", false, "Invert match (exclude)")
	fixed      = flag.Bool("F", false, "Fixed string matching (no patterns)")
	lineNum    = flag.Bool("n", false, "Print line numbers")
)

func main() {
	flag.Parse()

	// Получаем паттерн и файл из аргументов командной строки
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: grep [options] PATTERN FILE")
		return
	}
	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var re *regexp.Regexp
	if *fixed {
		re = regexp.MustCompile(regexp.QuoteMeta(pattern)) // Точное совпадение строки
	} else {
		if *ignoreCase {
			pattern = "(?i)" + pattern // Игнорировать регистр
		}
		re, err = regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("Error compiling pattern: %v\n", err)
			return
		}
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	matches := make([]int, 0)
	for i, line := range lines {
		match := re.MatchString(line)
		if *invert {
			match = !match // Инвертируем
		}
		if match {
			matches = append(matches, i)
		}
	}

	if *countOnly {
		fmt.Println(len(matches)) // Печатаем количество совпадений
		return
	}

	printMatches(lines, matches)
}

func printMatches(lines []string, matches []int) {
	printed := make(map[int]bool)
	for _, i := range matches {
		start := max(0, i-*before-*context)
		end := min(len(lines)-1, i+*after+*context)
		for j := start; j <= end; j++ {
			if _, alreadyPrinted := printed[j]; alreadyPrinted {
				continue
			}
			if *lineNum {
				fmt.Printf("%d:", j+1)
			}
			fmt.Println(lines[j])
			printed[j] = true
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
