package main

import (
	"fmt"
	"gopherSchool/exercises/fib"
	"net/http"
	"strconv"
)

type Implement struct {
	Num int
}

func (i Implement) MyMethod() int {
	return i.Num
}

type MyInterface interface {
	MyMethod() int
}

type MyStruct struct {
	i  MyInterface
	id int
}

const (
	a = "SDFSF"
	B
	b, c = iota, iota
)

func main() {

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func HttpFib() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		strNum := request.FormValue("num")
		intNum, err := strconv.Atoi(strNum)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		fibNum := fib.Fib(intNum)
		strNum = strconv.Itoa(fibNum)
		writer.WriteHeader(200)
		writer.Write([]byte(strNum))
	}
}
