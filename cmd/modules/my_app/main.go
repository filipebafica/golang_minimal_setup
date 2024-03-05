package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})

	fmt.Println("Server is listening on port !!8000...")
	http.ListenAndServe(":8000", nil)
}
