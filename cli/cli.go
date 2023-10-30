// Author: Moisés Adame Aguilar
// Creation Date: October 27th, 2023

package cli

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLI struct {}

// Constructor method for the Command Line Interface
func NewCLI() *CLI {
	return &CLI{}
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


func (cli *CLI) Run(){
	cli.validateArgs()

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	addBlockchainAddress := createBlockchainCmd.String("address", "", "Your addres.")

	setNodeIdCmd := flag.NewFlagSet("setid", flag.ExitOnError)
	idNumber := setNodeIdCmd.String("id", "", "Your node id.")

	logOutCmd := flag.NewFlagSet("logout", flag.ExitOnError)
	logOutValue := logOutCmd.String("value", "", "A bool true or false")

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "setid":
		err := setNodeIdCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "logout":
		err := logOutCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createBlockchainCmd.Parsed() {
		if *addBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*addBlockchainAddress)
	}

	if setNodeIdCmd.Parsed() {
		if *idNumber == "" {
			setNodeIdCmd.Usage()
			os.Exit(1)
		}
		cli.setID(*idNumber)
	}

	if logOutCmd.Parsed() {
		if *logOutValue == "" {
			logOutCmd.Usage()
			os.Exit(1)
		}
		cli.setID(*logOutValue)
	}
}