// Author: Moisés Adame Aguilar
// Date: May 11, 2023

package lib

// List of Blocks, The Blockchain
type Blockchain struct {
	Blocks []*Block
	size   int
}

// Constructor function for The Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{}, 0}
}

// Method that calculates the size of The Blockchain
func (bc *Blockchain) Size() int {
	return bc.size
}

// Method that adds a new block to The Blockchain
func (bc *Blockchain) AddBlock(data string) {
	if bc.Size() == 0 {
		bc.Blocks = append(bc.Blocks, NewBlock(data, nil))
	} else {
		var prevBlock *Block = bc.Blocks[bc.Size()-1]
		bc.Blocks = append(bc.Blocks, NewBlock(data, prevBlock.hash))
	}
	bc.size++
}
