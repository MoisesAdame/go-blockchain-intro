// Author: Moisés Adame Aguilar
// Creation Date: May 11th, 2023

package lib

import (
	"time"
	"encoding/hex"
	"fmt"
	"encoding/gob"
	"bytes"
	"log"
	"crypto/sha256"
)

// Main blockchain element, The Block
type Block struct {
	Timestamp     int64
	Nonce		  int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

// Method that geneterates the block's hash via PoW
func (block *Block) SetHash() {
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
}

// Constructor method for the Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), -1, transactions, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// Constructor method for Genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	block := &Block{time.Now().Unix(), 0, []*Transaction{coinbase}, []byte{}, []byte{}}
	block.SetHash()
	return block
}

// Printing the block's attributes.
func (block *Block) Print() {
	fmt.Println("[*] Timestamp: ", time.Unix(block.Timestamp, 0))
	fmt.Println("[*] Nonce: ", block.Nonce)

	prevHashBlockString := hex.EncodeToString(block.PrevBlockHash)
	if len(prevHashBlockString) != 0 {
		fmt.Println("[*] PrevBlockHash: ", prevHashBlockString[:4] + "..." + prevHashBlockString[len(prevHashBlockString) - 4:])
	}else{
		fmt.Println("[*] PrevBlockHash: <nil>")
	}
	
	hashString := hex.EncodeToString(block.Hash)
	fmt.Println("[*] Hash: ", hashString[:4] + "..." + hashString[len(hashString) - 4:])
}

// Serializing the block for its proper storage
func (block *Block) Encode() []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(*block)
	
	if err != nil {
		log.Fatal("encode:", err)
	}

	return buffer.Bytes()
}

// Deserializing the block to use it after its retrieval.
func DecodeBlock(buffer []byte) *Block {
	newBuffer := bytes.NewBuffer(buffer)
	dec := gob.NewDecoder(newBuffer)
	var block Block
	err := dec.Decode(&block)
	if err != nil {
		log.Fatal("decode:", err)
	}

	return &block
}

// Method that joins and hases every trabsaction id stored in the block.
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	
	hash := sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return hash[:]
}