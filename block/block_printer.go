package block

import (
	"fmt"

	"github.com/diegodisant/blockchain/common"
)

var _ common.Printable = (*Block)(nil)

func (block *Block) Print() {
	fmt.Printf("[-] Block: 0x%x\n", block.Hash)
	fmt.Printf("[+] Previous block: 0x%x\n", block.PrevHash)
	fmt.Printf("[+] PoW Number: %d\n", block.Pow)
	fmt.Printf("[+] Created at: %s\n", block.Timestamp)

	PrintBlockData(block.Data)

	fmt.Println()
}
