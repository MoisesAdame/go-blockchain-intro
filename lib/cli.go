// Author: Moisés Adame Aguilar
// Creation Date: October 27th, 2023

package lib

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLI struct {
	blockchain *Blockchain
}

// Constructor method for the Command Line Interface
func NewCLI() *CLI {
	return &CLI{NewBlockchain()}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddresses - Lists all addresses from the wallet file")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

// Mtethod that validates the input data.
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.blockchain.AddBlock(data)
}

func (cli *CLI) printChain() {
	cli.blockchain.Print()
}

func (cli *CLI) Run(){
	cli.validateArgs()

	createBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printBlockchainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := createBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := createBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createBlockCmd.Parsed() {
		if *addBlockData == "" {
			createBlockCmd.Usage()
			os.Exit(1)
		}
		fmt.Println("[*] Args:", *addBlockData)
		cli.addBlock(*addBlockData)
	}

	if printBlockchainCmd.Parsed() {
		cli.printChain()
	}
}