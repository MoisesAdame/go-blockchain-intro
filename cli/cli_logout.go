// Author: Moisés Adame Aguilar
// Creation Date: October 30th, 2023

package cli

import (
	"os"
	"fmt"
)

func (cli *CLI) logOut(value string){
	if value == "true"{
		os.Clearenv()
		fmt.Println("[*] Session successfully closed!")
	}else{
		fmt.Println("[*] Session still open!")
	}
}