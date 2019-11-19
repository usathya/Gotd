package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	var quotes [5]string
	quotes[0] = "40 is the old age of youth, while 50 is the youth of old age. - Victor Hugo"
	quotes[1] = "Knowledge is power. – Francis Bacon"
	quotes[2] = "Life is really simple, but we insist on making it complicated. - Confucius"
	quotes[3] = "This above all, to thine own self be true. - William Shakespeare"
	quotes[4] = "Live your dreams. – Les Brown"

	fmt.Fprintf(w, quotes[rand.Intn(5)])
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
