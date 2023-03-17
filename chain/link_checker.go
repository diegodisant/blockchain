package chain

import "github.com/diegodisant/blockchain/block"

type LinkChecker interface {
	CheckLink(prevBlock *block.Block, currentBlock *block.Block) bool
}

var _ LinkChecker = (*Blockchain)(nil)

func (blockchain *Blockchain) CheckLink(prevBlock *block.Block, currentBlock *block.Block) bool {
	return currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PrevHash != prevBlock.Hash
}
