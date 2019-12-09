package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	var quotes [6]string
	quotes[0] = "40 is the old age of youth, while 50 is the youth of old age. -- Victor Hugo"
	quotes[1] = "Knowledge is power. –- Francis Bacon"
	quotes[2] = "Life is really simple, but we insist on making it complicated. -- Confucius"
	quotes[3] = "This above all, to thine own self be true. -- William Shakespeare"
	quotes[4] = "Live your dreams. -– Les Brown"
	quotes[5] = "Doo be doo be dooo. -- Frank Sinatra"

	fmt.Fprintf(w, quotes[rand.Intn(6)])
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Println("QOTD listening on port 10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
