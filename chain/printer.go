package chain

import (
	"fmt"

	"github.com/diegodisant/blockchain/common"
)

var _ common.Printable = (*Blockchain)(nil)

func (blockchain *Blockchain) Print() {
	fmt.Println("[-] Blockchain")
	fmt.Printf("[+] Difficulty: %d\n\n", blockchain.Difficulty)

	for _, block := range blockchain.Chain {
		block.Print()
	}
}
