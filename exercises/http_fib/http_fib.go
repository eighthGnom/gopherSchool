package main

import (
	"gopherSchool/exercises/fib"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/fibonachi", HttpFib())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
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
