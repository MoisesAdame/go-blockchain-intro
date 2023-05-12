// Author: Moisés Adame Aguilar
// Date: May 11, 2023

package lib

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// Main blockchain element, The Block
type Block struct {
	timestamp int64
	nonce     int64
	data      []byte
	prevHash  []byte
	hash      []byte
}

// Constructor method for the Block
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), -1, []byte(data), prevHash, []byte{}}

	// Mining (obtaining nonce and hash)
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	// Asigning the mined values
	block.nonce = int64(nonce)
	block.hash = hash

	return block
}

// Function that converts an int into a []byte of an hex
func IntToHex(num int64) []byte {
	return []byte(strconv.FormatInt(num, 16))
}

// Printing the block's attributes.
func (b *Block) Print() {
	// Pasing hashes from []byte to string
	var hexPrevHash string = hex.EncodeToString(b.prevHash)
	var hexHash string = hex.EncodeToString(b.hash)

	// Main string
	var txt string
	txt += "- Timestamp: " + strconv.FormatInt(b.timestamp, 10) + "\n"
	txt += "- Data:      " + string(b.data) + "\n"
	txt += "- Nonce:     " + strconv.FormatInt(b.nonce, 10) + "\n"

	if len(hexPrevHash) == 0 {
		txt += "- Prev Hash: " + "<nil>" + "\n"
	} else {
		txt += "- Prev Hash: " + hexPrevHash[:4] + "..." + hexPrevHash[len(hexPrevHash)-4:] + "\n"
	}
	txt += "- Hash:      " + hexHash[:4] + "..." + hexHash[len(hexHash)-4:] + "\n"

	fmt.Println(txt)
}

// Serializing the block for its proper storage
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encode := gob.NewEncoder(&res)
	encode.Encode(b)
	return res.Bytes()
}

// Deserializing the block to use it after its retrieval.
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)

	return &block
}
