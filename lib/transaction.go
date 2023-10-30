// Author: Moisés Adame Aguilar
// Creation Date: October 29th, 2023

package lib

import(
	"fmt"
	"encoding/gob"
	"crypto/sha256"
	"log"
	"bytes"
)

const subsidy = 10

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := NewTXInput([]byte{}, -1, data)
	txout := NewTXOutput(subsidy, to)
	tx := &Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.setHash()

	return tx
}

// Serializing the transaction for its proper hashing.
func (tx *Transaction) encode() []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(*tx)
	
	if err != nil {
		log.Fatal("[*] Transaction encoding error: ", err)
	}

	return buffer.Bytes()
}

func (tx *Transaction) setHash() {
	hash := sha256.Sum256(tx.encode())
	tx.ID = hash[:]
}