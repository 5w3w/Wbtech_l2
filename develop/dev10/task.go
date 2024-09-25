package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	timeoutArg := flag.Duration("timeout", 10*time.Second, "Timeout for the connection (default: 10s)")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Установка таймера
	conn, err := net.DialTimeout("tcp", address, *timeoutArg)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", address)

	// Запуск горутины для чтения данных из соединения и вывода их в STDOUT
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {

			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Connection closed:", err)
		}
		os.Exit(0)
	}()

	// Чтение данных из STDIN и отправка их в соединение
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := fmt.Fprintln(conn, scanner.Text())
		if err != nil {
			fmt.Println("Failed to write to connection:", err)
			return
		}
	}

	// Проверка ошибки на закрытие сокета
	if scanner.Err() != nil {
		fmt.Println("Error reading from stdin:", scanner.Err())
	}

	fmt.Println("Closing connection.")
}
