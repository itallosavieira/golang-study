package main

type ID int // personal type

var (
	b bool
	c int = 30
	e ID  = 1
)

const d float64 = 1.2

func main() {
	a := "Var shorthand"
	println(a, b, c, d, e)
}
