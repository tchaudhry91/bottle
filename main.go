package main

import (
	"fmt"
	"github.com/tchaudhry91/bottle/bottle"
	"os"
)

func main() {
	b := bottle.NewBottle("try-me")
	b.Fill(os.Stdin)

	fmt.Println("Pouring..")
	err := b.Pour(os.Stdout)
	if err != nil {
		panic(err)
	}

	fmt.Println("Draining..")
	err = b.Drain(os.Stdout)
	if err != nil {
		panic(err)
	}

}
