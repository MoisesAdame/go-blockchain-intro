package main

import (
	"./lib"
)

func main() {
	blockchain := lib.NewBlockchain()
	blockchain.AddBlock("Bloque 1")
	blockchain.AddBlock("Bloque 2")
	blockchain.AddBlock("Bloque 3")
	blockchain.Print()
}