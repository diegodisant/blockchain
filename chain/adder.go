package chain

import (
	"time"

	"github.com/diegodisant/blockchain/block"
	"github.com/diegodisant/blockchain/common"
)

var _ common.Appendable = (*Blockchain)(nil)

func (blockchain *Blockchain) Append(from, to string, amount float64) {
	var blockData = block.BuildBlockData(from, to, amount)

	var chainSize = len(blockchain.Chain)
	var lastBlock = blockchain.Chain[chainSize-1]

	var newBlock = &block.Block{
		Data:      blockData,
		PrevHash:  lastBlock.Hash,
		Timestamp: time.Now(),
	}

	newBlock.Mine(blockchain.Difficulty)

	blockchain.Chain = append(blockchain.Chain, newBlock)
}
