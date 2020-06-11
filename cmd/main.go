package main

import (
	"fmt"
	"math/rand"
	"time"

	gen "github.com/ustaty/ustatynamegenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(gen.GetRandomName(0))
}
