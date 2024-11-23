package main

import "fmt"

type Printable interface {
	~int | ~float64
	String() string
}

type PrintableInt int

func (p PrintableInt) String() string {
	return fmt.Sprintf("%d", p)
}

type PrintableFloat64 float64

func (p PrintableFloat64) String() string {
	return fmt.Sprintf("%f", p)
}

func Print[T Printable](a T) {
	fmt.Println(a)
}

func main() {
	var myInt PrintableInt
	myInt = 55

	Print(myInt)
	// fmt.Println(myInt.String())

	var myFloat PrintableFloat64
	myFloat = 55.5

	fmt.Println(myFloat.String())
}
