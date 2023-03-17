package block

import (
	"strings"

	"github.com/diegodisant/blockchain/common"
)

var _ common.Minable = (*Block)(nil)

func (block *Block) Mine(difficulty int) {
	var zeroesPrefix = strings.Repeat("0", difficulty)

	for !strings.HasPrefix(block.Hash, zeroesPrefix) {
		block.Pow += 1
		block.Hash = block.CalculateHash()
	}
}
