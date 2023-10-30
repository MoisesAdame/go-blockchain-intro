// Author: Moisés Adame Aguilar
// Creation Date: October 29th, 2023

package lib

import(

)

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func NewTXOutput(value int, address string) TXOutput {
	return TXOutput{value, address}
}