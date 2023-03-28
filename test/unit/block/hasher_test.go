package block_test

import (
	"testing"
	"time"

	"github.com/diegodisant/blockchain/block"
	"github.com/diegodisant/blockchain/test"
)

func TestCalculateHash(t *testing.T) {
	var dataProvider = buildCalculateHashDataProvider()

	for providerName, providerData := range dataProvider {
		t.Logf("Running provider: (%s)", providerName)

		currentBlock, canCastBlock := providerData.(*block.Block)

		if !canCastBlock {
			t.Error("Cannot current block")
		}

		currentBlock.Hash = currentBlock.CalculateHash()
		var currentHash = currentBlock.CalculateHash()

		t.Logf("Block hash: %x", currentBlock.Hash)
		t.Logf("Generated hash: %x", currentHash)

		if currentBlock.Hash != currentHash {
			t.Errorf(
				"Generated hash for same block are different with PoW: %d",
				currentBlock.Pow,
			)
		}
	}
}

func buildCalculateHashDataProvider() test.DataProvider {
	return test.DataProvider{
		"calculate hash with correct block data": &block.Block{
			Data:      block.BuildBlockData("sender", "receiver", 600.32),
			PrevHash:  "0x00",
			Timestamp: time.Now(),
		},
		"calculate hash with corrupt block data": &block.Block{
			Data:      block.BuildBlockData("sender", "receiver\"'''}}}\\}", 599.99),
			PrevHash:  "0x00",
			Timestamp: time.Now().Add(time.Minute),
		},
	}
}
