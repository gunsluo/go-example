package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Now server is running on port 3000.")
	http.ListenAndServe(":3000", nil)
}
