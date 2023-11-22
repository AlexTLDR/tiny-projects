// run the program by adding the flag
// go run helloworld.go -lang=en
// check localhost:8080

package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	var lang string
	flag.StringVar(&lang, "lang", "ro", "The required language, e.g. en, ro...")
	flag.Parse()
	greeting := greet(language(lang))
	fmt.Fprintf(writer, "%s, %s!", greeting, request.URL.Path[1:])
}

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
	"ro": "Salut lume",        // Romanian
}

// greet returns a greeting to the world.
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
