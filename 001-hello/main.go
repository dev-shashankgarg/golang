package main

import (
	"flag"
	"fmt"
)

type language string

var phrases = map[language]string{
	"en": "Hello World!",
	"fr": "Bonjour le monde!",
	"hi": "नमस्कार, संसार!",
}

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "Language of Choice such as en, fr, hi ..etc")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(l language) string {
	greeting, ok := phrases[l]
	if !ok {
		return fmt.Sprintf("Unsupported Language: %q", l)
	}
	return greeting
}
