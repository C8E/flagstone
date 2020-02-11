package main

import (
	"flag"
	"fmt"

	"github.com/C8E/flagstone"
)

var who *string

func main() {
	who = flag.String("who", "world", "say hello to ...")
	flag.Parse()

	flagstone.Launch("helloworld", "flagstone sample")

	fmt.Println("hello " + *who + "!")
}
