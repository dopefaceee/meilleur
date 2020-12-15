package main

import (
	"fmt"

	"github.com/meilleur/rand"
)

func main() {
	fmt.Println(rand.String(20))
	fmt.Println(rand.RememberToken())
}
