/*
Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.

Применимость:
- Когда вы хотите параметризовать объекты выполняемым действием.
- Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
- Когда вам нужна операция отмены.
Плюсы:
- позволяет откладывать выполнение команд, выстраивать их в очереди, а также хранить историю и делать отмену.
- Убирает прямую зависимость между объектами, вызываю-
щими операции, и объектами, которые их непосредственно
выполняют.
- Позволяет реализовать простую отмену и повтор операций.
- Позволяет реализовать отложенный запуск операций.
- Позволяет собирать сложные команды из простых.
- Реализует принцип открытости/закрытости.
Минусы: Усложняет код программы из-за введения множества допол-
нительных классов
Реальные примеры
*/

package main

import "fmt"

type Command interface {
	execute()
}

type Device interface {
	on()
	off()
}

type Button struct {
	command Command
}

type OnCommand struct {
	device Device
}

type OffCommand struct {
	device Device
}

type Tv struct {
	isRunning bool
}

func (b *Button) press() {
	b.command.execute()
}

func (c *OnCommand) execute() {
	c.device.off()
}

func (c *OffCommand) execute() {
	c.device.on()
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

}
