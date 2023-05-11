package main

import (
	"./lib"
	"fmt"
	"strconv"
)

func main() {
	var blockchain *lib.Blockchain = lib.NewBlockchain()
	blockchain.AddBlock("block1")
	blockchain.AddBlock("block2")

	for index, val := range blockchain.Blocks {
		fmt.Println("[*] Block " + string(index))
		val.Print()
		pow := lib.NewProofOfWork(val)
		fmt.Println("- POW: ", strconv.FormatBool(pow.Validate()))
	}
}
