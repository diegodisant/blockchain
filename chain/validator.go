package chain

import "github.com/diegodisant/blockchain/common"

var _ common.Validable = (*Blockchain)(nil)

func (blockchain *Blockchain) IsValid() bool {
	for blockIndex := range blockchain.Chain[1:] {
		prevBlock := blockchain.Chain[blockIndex]
		currentBlock := blockchain.Chain[blockIndex+1]

		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.Hash == prevBlock.Hash {
			return false
		}
	}

	return true
}
