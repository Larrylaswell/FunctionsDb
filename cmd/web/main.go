package main

import (
	"fmt"
	"net/http"
	"handlers"
)

const PortNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting app on port %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
