package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("asdf")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, "xatta")

		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("bytes written %d ", n))
	})

	_ = http.ListenAndServe(":3000", nil)
}
