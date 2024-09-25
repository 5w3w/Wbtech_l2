package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Custom shell. Type \\quit to exit.")

	for {
		fmt.Print("shell> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "\\quit" || input == "\\q" {
			break
		}

		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			pipeCommands(commands)
		} else {
			executeCommand(strings.TrimSpace(input))
		}
	}
}

// Функция для обработки команд
func executeCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "cd":
		if len(parts) > 1 {
			err := os.Chdir(parts[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("cd: missing argument")
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(dir)
		}
	case "echo":
		fmt.Println(strings.Join(parts[1:], " "))
	case "kill":
		if len(parts) > 1 {
			pid, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				err := syscall.Kill(pid, syscall.SIGKILL)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		} else {
			fmt.Println("kill: missing argument")
		}
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		runCommand(parts)
	}
}

// Запуск внешних команд с помощью exec
func runCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Обработка пайпов
func pipeCommands(commands []string) {
	var lastCmd *exec.Cmd
	for i, cmdStr := range commands {
		parts := strings.Fields(cmdStr)
		cmd := exec.Command(parts[0], parts[1:]...)

		if i == 0 {
			lastCmd = cmd
		} else {
			out, err := lastCmd.StdoutPipe()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cmd.Stdin = out
		}

		if i == len(commands)-1 {
			cmd.Stdout = os.Stdout
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("Error:", err)
			return
		}

		lastCmd = cmd
	}

	if lastCmd != nil {
		lastCmd.Wait()
	}
}
