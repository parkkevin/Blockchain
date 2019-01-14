package main

import (
	"blockchain"
	"fmt"
)

func main() {

	b0 := blockchain.Initial(7)
	b0.Mine(1)
	b := blockchain.Blockchain{}
	b.Add(b0)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b.Add(b1)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	b.Add(b2)
	if b.IsValid() {
		fmt.Printf("The blockchain is valid: %t\n", true)
	} else {
		fmt.Printf("The blockchain in valid: %t\n", false)
	}
}
