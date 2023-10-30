// Author: Moisés Adame Aguilar
// Creation Date: October 29th, 2023

package lib

import(

)

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

func NewTXInput(txid []byte, vout int, scriptSig string) TXInput {
	return TXInput{txid, vout, scriptSig}
}