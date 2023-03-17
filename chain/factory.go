package chain

import (
	"time"

	"github.com/diegodisant/blockchain/block"
)

func BuildBlockchain(difficulty int) *Blockchain {
	genesisBlock := &block.Block{
		Hash:      "0",
		Timestamp: time.Now(),
	}

	return &Blockchain{
		GenesisBlock: genesisBlock,
		Chain:        []*block.Block{genesisBlock},
		Difficulty:   difficulty,
	}
}
