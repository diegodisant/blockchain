package main

import (
	"fmt"

	"github.com/diegodisant/blockchain/chain"
)

func main() {
	var blockchain = chain.BuildBlockchain(2)

	blockchain.Append("Juni", "Yiyi", 5000.32)
	blockchain.Append("Pillu", "Sata", 35000.28)

	blockchain.Print()

	fmt.Printf("Is this blockchain valid? %v\n", blockchain.IsValid())
}
