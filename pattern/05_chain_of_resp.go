/*
Цепочка обязанностей — это поведенческий паттерн
проектирования, который позволяет передавать запросы
последовательно по цепочке обработчиков. Каждый
последующий обработчик решает, может ли он обработать
запрос сам и стоит ли передавать запрос дальше по цепи.

Применимость: Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,260
Поведенческие паттерны / Цепочка обязанностей какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
Плюсы:
- Избавляет от жёсткой привязки отправителя запроса к его получателю, позволяя выстраивать цепь из различных обработчиков динамически.
- Уменьшает зависимость между клиентом и обработчиками.
- Реализует принцип единственной обязанности.
- Реализует принцип открытости/закрытости.
Минусы:
- Запрос может остаться никем не обработанным.
Реальные примеры
*/
package main

import "fmt"

type Departament interface {
	execute(*Patient)
	setNext(Departament)
}

type Reception struct {
	next Departament
}

type Doctor struct {
	next Departament
}

type Medical struct {
	next Departament
}

type Cashier struct {
	next Departament
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Departament) {
	r.next = next
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Departament) {
	d.next = next
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Departament) {
	m.next = next
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment done")

	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) setNext(next Departament) {
	c.next = next
}

func main() {

	cashier := &Cashier{}
	medical := &Medical{}
	doctor := &Doctor{}
	reception := &Reception{}
	patient := &Patient{name: "abc"}
	medical.setNext(cashier)
	doctor.setNext(medical)
	reception.setNext(doctor)
	reception.execute(patient)
}
