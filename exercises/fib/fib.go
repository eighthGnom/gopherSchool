package fib

var dataMap map[int]int

func init() {
	dataMap = make(map[int]int)
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	if fib, ok := dataMap[n]; ok {
		return fib
	}
	fib := Fib(n-1) + Fib(n-2)
	dataMap[n] = fib
	return fib
}
