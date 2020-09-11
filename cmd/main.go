package main

import (
	"fmt"

	"github.com/drrrMikado/shorten/rand"
)

func main() {
	s := rand.String(8)
	fmt.Println(s)
}
