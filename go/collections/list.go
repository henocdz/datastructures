package collections

type List interface {
	IsEmpty() bool
	Push(data int)
	Pop() int
	Print()
}
