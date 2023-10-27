// Author: Moisés Adame Aguilar
// Creation Date: May 11th, 2023

package lib

import (
	"math/big"
	"bytes"
	"strconv"
	"math"
	"crypto/sha256"
)

// Difficulty, target
const targetBits = 24

// Proof of Work using Hascash
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// Constructor of PoW for a given block
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	return &ProofOfWork{block, target}
}

// Prepare the contents to be hashed.
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
			[][]byte{
				[]byte(strconv.FormatInt(pow.block.Timestamp, 10)),
				pow.block.Data, 
				pow.block.PrevBlockHash,
				[]byte(strconv.FormatInt(targetBits, 10)),
				[]byte(strconv.FormatInt(nonce, 10)),
			}, []byte{})

	return data
}

// Running the PoW algorithm
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hash [32]byte
	var hashInt big.Int
	var nonce int64 = 0
	var maxNonce int64 = math.MaxInt64

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return int64(nonce), hash[:]
}

// Validate the proof of work using the bolck's nonce
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var hash [32]byte
	data := pow.prepareData(pow.block.Nonce)
	hash = sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}