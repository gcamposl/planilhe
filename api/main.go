package main

import (
	"api/src"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("init")

	r := src.Generate()
	fmt.Println(r)

	log.Fatal(http.ListenAndServe(":5000", r))
}
