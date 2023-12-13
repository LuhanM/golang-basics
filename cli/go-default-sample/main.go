package main

import (
	"flag"
	"fmt"
)

func main() {
	channel := flag.String("name", "Luhan Meireles", "Seu nome completo")
	flag.Parse()
	fmt.Println(*channel)
}
