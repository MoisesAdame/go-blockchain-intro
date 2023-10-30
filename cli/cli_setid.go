// Author: Moisés Adame Aguilar
// Creation Date: October 30th, 2023

package cli

import (
	"os"
	"fmt"
)

func (cli *CLI) setID(nodeID string){
	os.Setenv("NODE_ID", nodeID)
	fmt.Println("[*] NODE_ID set as: ", nodeID)
}