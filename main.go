package main

import (
	"./lib"


)

func main() {
	var blockchain *lib.Blockchain = lib.NewBlockchain()
	blockchain.AddBlock("block1")
	blockchain.AddBlock("block2")
}
