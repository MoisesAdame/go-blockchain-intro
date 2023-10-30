// Author: Moisés Adame Aguilar
// Creation Date: May 11th, 2023

package lib

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const database = "blockchain_%s.db"
const blocksBucket = "blocks"

// List of Blocks, The Blockchain
type Blockchain struct {
	Database *bolt.DB
	Head 	 []byte
}

// Method that instatiates a totally new blockchain for some given address and nodeID.
func CreateBlockchain(address, nodeID string) *Blockchain {
	var currentHead []byte
	nodeFileName := fmt.Sprintf(database, nodeID)

	if dbExists(nodeFileName) {
		log.Fatal("[*] Blockchain already exists!")
		return nil
	}else{
		db, err := bolt.Open(nodeFileName, 0600, nil)

		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			bucket, bucketError := tx.CreateBucket([]byte(blocksBucket))

			cbtx := NewCoinbaseTX(address, "")
			genesisBlock := NewGenesisBlock(cbtx)

			bucketError = bucket.Put([]byte(genesisBlock.Hash), genesisBlock.Encode())
			bucketError = bucket.Put([]byte("l"), genesisBlock.Hash)
			currentHead = genesisBlock.Hash
			
			return bucketError
		})

		if err != nil {
			log.Fatal("[*] Error building blockchain: ", err)
		}

		blockchain := &Blockchain{db, currentHead}
		return blockchain
	}
}

// Constructor function for The Blockchain.
func NewBlockchain(nodeID string) *Blockchain {
	var currentHead []byte
	nodeFileName := fmt.Sprintf(database, nodeID)

	if !dbExists(nodeFileName) {
		log.Fatal("[*] No existing blockchain found. Create one first!")
		return nil
	}else{
		db, err := bolt.Open(nodeFileName, 0600, nil)

		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(blocksBucket))

			currentHead = bucket.Get([]byte("l"))
			return nil
		})

		if err != nil {
			log.Fatal("[*] Error building blockchain: ", err)
		}

		blockchain := &Blockchain{db, currentHead}
		return blockchain
	}
}

// Method that adds a new block to The Blockchain
func (blockchain *Blockchain) AddBlock(newBlock *Block) {
	db, err := bolt.Open(database, 0600, nil)
	if err != nil {
		log.Fatal("Unable to open database!")
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		bucket.Put([]byte(newBlock.Hash), newBlock.Encode())
		bucket.Put([]byte("l"), newBlock.Hash)
		blockchain.Head = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Fatal("Unable create Genesis Block!")
	}
}

// Method that prints every block within the chain.
func (blockchain *Blockchain) Print() {
	db, err := bolt.Open(database, 0600, nil)
	if err != nil {
		log.Fatal("Unable to open database!")
	}

	defer db.Close()

	auxHead := blockchain.Head
	var currentBlock *Block
	
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		for len(auxHead) != 0 {
			currentBlock = DecodeBlock(bucket.Get(auxHead))
			currentBlock.Print()

			pow := NewProofOfWork(currentBlock)
			fmt.Println("[*] PoW: ", pow.Validate())
			fmt.Println()

			auxHead = currentBlock.PrevBlockHash
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error reading database!")
	}
}

// Function that checks if the database file exists.
func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}