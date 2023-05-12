// Author: Moisés Adame Aguilar
// Date: May 11, 2023

package lib

import (
	"github.com/boltdb/bolt"
)

const database = "blockchain.db"
const blocksBucket = "blocks"

// List of Blocks, The Blockchain
type Blockchain struct {
	tip  []byte
	db   *bolt.DB
	size int
}

// Constructor function for The Blockchain
func NewBlockchain() *Blockchain {
	// Creating the tip, the head of our blockchain
	var tip []byte

	// Opening the database
	db, _ := bolt.Open(database, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		// Cheching for trhe existance of the bucket
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			// Instantiating the genesis block
			var genesisBlock *Block = NewBlock("Genesis", nil)

			// Serializing and storing
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			b.Put(genesisBlock.hash, genesisBlock.Serialize())
			b.Put([]byte("l"), genesisBlock.hash)

			// Making the tip the genesis block
			tip = genesisBlock.hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	return &Blockchain{tip, db, 1}
}

// Method that calculates the size of The Blockchain
func (bc *Blockchain) Size() int {
	return bc.size
}

// Method that adds a new block to The Blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	newBlock := NewBlock(data, lastHash)
	newBlock.Print()

	bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		b.Put(newBlock.hash, newBlock.Serialize())
		b.Put([]byte("l"), newBlock.hash)
		bc.tip = newBlock.hash

		return nil
	})

	bc.size++
}
