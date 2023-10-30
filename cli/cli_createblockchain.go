// Author: Moisés Adame Aguilar
// Creation Date: October 30th, 2023

package cli

import (
	"../lib"
	"os"
	"fmt"
)

func (cli *CLI) createBlockchain(address string){

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Println("[*] NODE_ID env. var is not set!")
		os.Exit(1)
	}

	lib.CreateBlockchain(address, nodeID)
}