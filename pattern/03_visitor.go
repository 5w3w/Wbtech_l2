/*

Паттерн Посетитель -  это поведенческий паттерн, который позволяет
добавить новую операцию для целой иерархии классов, не изменяя код этих классов.

Применимость:
- Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции, но вы не хотите «засорять» классы такими операциями.
- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
Плюсы:
- Упрощает добавление операций, работающих со сложными
структурами объектов.
- Объединяет родственные операции в одном классе.
- Посетитель может накапливать состояние при обходе структуры элементов
Минусы:
- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.
Реальные примеры

*/

package main

import "fmt"

const PI = 3.14

type Shape interface {
	getType() string
	accept(Visitor)
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Square struct {
	side float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius float64
}

func (c *Circle) getType() string {
	return "Circle"
}

func (s *Circle) accept(v Visitor) {
	v.visitForCircle(s)
}

type Rectangle struct {
	a float64
	b float64
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType(v Visitor) string {
	return "Rectangle"
}

type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitForCircle(c *Circle) {

	a.area = c.radius * c.radius * PI
	fmt.Println("Расчет площади для круга, S = ", a.area)

}

func (a *AreaCalculator) visitForSquare(s *Square) {
	a.area = s.side * s.side
	fmt.Println("Расчет площади для квадрата S =", a.area)
}

func (a *AreaCalculator) visitForRectangle(t *Rectangle) {
	a.area = t.a * t.b
	fmt.Println("Расчет площади для треугольника S =", a.area)
}

type MiddleCoordinates struct {
	x float64
	y float64
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Расчет координатов средней точки для круга")
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Расчет координатов средней точки для квадрата")
}
func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Расчет координатов средней точки для треугольника")

}

func main() {
	square := &Square{side: 3}
	circle := &Circle{radius: 5}
	Rectangle := &Rectangle{a: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	Rectangle.accept(areaCalculator)

	fmt.Println()

	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	Rectangle.accept(middleCoordinates)

}
