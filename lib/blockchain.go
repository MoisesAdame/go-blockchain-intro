// Author: Moisés Adame Aguilar
// Creation Date: May 11th, 2023

package lib

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const database = "blockchain.db"
const blocksBucket = "blocks"

// List of Blocks, The Blockchain
type Blockchain struct {
	Database *bolt.DB
	Head 	 []byte
}

// Constructor function for The Blockchain
func NewBlockchain() *Blockchain {
	var currentHead []byte
	db, err := bolt.Open(database, 0600, nil)
	if err != nil {
		log.Fatal("Unable to create database!")
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			bucket, bucketError := tx.CreateBucket([]byte(blocksBucket))

			genesisBlock := NewGenesisBlock()

			bucketError = bucket.Put([]byte(genesisBlock.Hash), genesisBlock.Encode())
			bucketError = bucket.Put([]byte("l"), genesisBlock.Hash)
			currentHead = genesisBlock.Hash

			return bucketError
		}else{
			currentHead = bucket.Get([]byte("l"))
			return nil
		}
	})

	if err != nil {
		log.Fatal("Unable create Genesis Block!")
	}

	blockchain := &Blockchain{db, currentHead}
	return blockchain
}

// Method that adds a new block to The Blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	db, err := bolt.Open(database, 0600, nil)
	if err != nil {
		log.Fatal("Unable to open database!")
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		prevBlockHash := bucket.Get([]byte("l"))
		newBlock := NewBlock(data, prevBlockHash)

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