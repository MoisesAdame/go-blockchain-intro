// Author: Moisés Adame Aguilar
// Date: May 11, 2023

package lib

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

// Difficulty, target
const targetBits = 24

// Proof of Work using Hascash
type ProofOfWork struct {
	block  *Block
	Target *big.Int
}

// Constructor of POW for a given block
func NewProofOfWork(b *Block) *ProofOfWork {
	Target := big.NewInt(1)
	Target.Lsh(Target, uint(256-targetBits))
	return &ProofOfWork{b, Target}
}

// Prepare the contents to be hashed.
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{IntToHex(pow.block.timestamp),
			pow.block.data,
			pow.block.prevHash,
			IntToHex(targetBits),
			IntToHex(int64(nonce))}, []byte{})

	return data
}

// Running the POW algorithm
func (pow *ProofOfWork) Run() (int, []byte) {
	var preparedData []byte
	var hashedData [32]byte
	var hashInt big.Int
	var nonce int = 0

	for true {
		// Prepare
		preparedData = pow.PrepareData(nonce)

		// Hash the prepared data
		hashedData = sha256.Sum256(preparedData)

		// Convert and compare
		hashInt.SetBytes(hashedData[:])
		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hashedData[:]
}

// Validate the proof of work using the bolck's nonce
func (pow *ProofOfWork) Validate() bool {
	// Prepare
	var preparedData []byte = pow.PrepareData(int(pow.block.nonce))

	// Hash 
	var hash [32]byte = sha256.Sum256(preparedData)

	// Move to Big-Int
	var hashInt big.Int
	hashInt = *hashInt.SetBytes(hash[:])

	// Compare
	return hashInt.Cmp(pow.Target) == -1
}