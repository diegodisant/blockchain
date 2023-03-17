package chain

import "github.com/diegodisant/blockchain/block"

type Blockchain struct {
	GenesisBlock *block.Block
	Chain        []*block.Block
	Difficulty   int
}
