package hello

import (
	"log"

	"golang.org/x/example/stringutil"
)

func Hello() string {
	return stringutil.Reverse("Hello")
}

func main() {
	log.Println(Hello())
}
