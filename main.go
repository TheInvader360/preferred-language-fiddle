package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/croeso", croesoHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	fmt.Printf("Example usage:\n")
	fmt.Printf("* request http://localhost:%d in a suitably configured browser\n", port)
	fmt.Printf("* curl -s -H 'Accept-Language: cy,en;q=0.9,fr;q=0.8,de;q=0.7' http://localhost:%d/check\n\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
