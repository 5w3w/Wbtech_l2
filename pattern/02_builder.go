/*

Паттерн строитель - это порождающий паттерн проектирования, который позволяет создавать объекты пошагово.

Применимость: Паттерн Строитель предлагает вынести
конструирование объекта за пределы его собственного класса, поручив это дело отдельным объектам,
которые следует называть строителями.

Плюсы: Позволяет пошагово создавать продукт зависящий, от маленьких частей
Позволяет использовать один и тот же код для одних объектов
Изолирует сложный код объекта и его основной бизнес логики.
Минусы: Усложняет код из за введения дополнительных классов/структур/интерфейсов
клиент привязан к конкретному объекту строителя и если не будет какого-то метода, разработчику нужно будет его добавить
Реальные примеры

*/

package main

import "fmt"

type Car struct {
	Color Color
	Speed Speed
	Wheel Wheel
}

type Builder interface {
	setColor()
	setWheels()
	setTopSpeed()
	BuildCar() Car
}

type NormalBuilder struct {
	Color Color
	Speed Speed
	Wheel Wheel
}

type AtBuilder struct {
	Color Color
	Speed Speed
	Wheel Wheel
}

// type Interface interface{
// 	Drive() error
// 	Stop() error
// }

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60
)

type Color string

const (
	Bluecolor  Color = "blue"
	Greencolor       = "Green"
	RedColor         = "Red"
)

type Wheel string

const (
	SportsWheels = "Sport"
	SteelWheels  = "steel"
)

func getBuilder(builderType string) Builder {
	if builderType == "NormalBuilder" {
		return newNormalbuilder()
	}
	if builderType == "AtBuilder" {
		return newAtBuilder()
	}
	return nil

}

func newNormalbuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setColor() {
	b.Color = Bluecolor
}

func (b *NormalBuilder) setWheels() {
	b.Wheel = SteelWheels
}

func (b *NormalBuilder) setTopSpeed() {
	b.Speed = KPH * 60
}

func (b *NormalBuilder) BuildCar() Car {
	return Car{
		Color: b.Color,
		Wheel: b.Wheel,
		Speed: b.Speed,
	}
}

func newAtBuilder() *AtBuilder {
	return &AtBuilder{}
}

func (n *AtBuilder) setColor() {
	n.Color = RedColor
}

func (n *AtBuilder) setWheels() {
	n.Wheel = SportsWheels
}

func (n *AtBuilder) setTopSpeed() {
	n.Speed = MPH * 90
}
func (n *AtBuilder) BuildCar() Car {
	return Car{
		Color: n.Color,
		Wheel: n.Wheel,
		Speed: n.Speed,
	}
}

type Direcotor struct {
	builder Builder
}

func newDirector(b Builder) *Direcotor {
	return &Direcotor{
		builder: b,
	}
}
func (d *Direcotor) setBuilder(b Builder) {
	d.builder = b
}

func (d *Direcotor) BuildCar() Car {
	d.builder.setColor()
	d.builder.setWheels()
	d.builder.setTopSpeed()
	return d.builder.BuildCar()

}

func main() {

	NormalBuilder := getBuilder("NormalBuilder")
	AtCarBuilder := getBuilder("AtBuilder")

	director := newDirector(NormalBuilder)
	normalCar := director.BuildCar()

	fmt.Printf("Normal Car Color: %s\n", normalCar.Color)
	fmt.Printf("Normal Car Wheel: %s\n", normalCar.Wheel)
	fmt.Printf("Normal Car Speed: %s\n", normalCar.Speed)

	director.setBuilder(AtCarBuilder)
	AtCar := director.BuildCar()
	fmt.Printf("AtCar Car Color: %s\n", AtCar.Color)
	fmt.Printf("AtCar Car Wheel: %s\n", AtCar.Wheel)
	fmt.Printf("AtCar Car Speed: %s\n", AtCar.Speed)

}
