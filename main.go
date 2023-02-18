package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	// game
	runtime.GOMAXPROCS(runtime.NumCPU())
	handlerFn := func(writer http.ResponseWriter, request *http.Request) {
		fmt.Print(request.Header)
		writer.Write([]byte("hello"))
	}
	http.HandleFunc("/", handlerFn)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello, game")
}
