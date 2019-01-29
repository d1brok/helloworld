package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Got request\n")
		time.Sleep(200 * time.Millisecond)
		text := "Dogs Rule!!! Cats drool!!!"
		fmt.Fprintf(w, os.Getenv("K_REVISION")+": "+text+"\n")
	})

	fmt.Print("Listening on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
