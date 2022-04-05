package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	seed, _ := strconv.Atoi(os.Args[1])
	count, _ := strconv.Atoi(os.Args[2])
	rand.Seed(int64(seed))
	for i := count; i > 0; i-- {

		fmt.Println(rand.Uint64())
	}
}
