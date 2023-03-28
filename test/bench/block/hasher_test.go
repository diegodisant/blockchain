package hasher_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/diegodisant/blockchain/block"
)

func BenchmarkCalculateHash(b *testing.B) {
	b.Log("Calculate 1 million times the same block hash")

	var currentBlock = &block.Block{
		PrevHash:  "0",
		Data:      block.BuildBlockData("Alex", "Simon", 6000),
		Timestamp: time.Now(),
	}

	for currentIteration := 1; currentIteration <= 1e6; currentIteration += 1 {
		currentBlock.Hash = currentBlock.CalculateHash()
	}
}

func BenchmarkCalculateHashWithChangingData(b *testing.B) {
	b.Log("Calculate 1 million times the block hash with changing the data")

	var currentBlock = &block.Block{
		PrevHash:  "0",
		Data:      block.BuildBlockData("Alex", "Simon", 6000),
		Timestamp: time.Now(),
	}

	var currentHash string

	for currentIteration := 1; currentIteration <= 1e6; currentIteration += 1 {
		currentBlock.Data["amount"] = rand.Float64() * 1000

		currentHash = currentBlock.CalculateHash()

		if currentHash == currentBlock.Hash {
			b.Errorf("Block data is different but the hash didn't change: %s", currentHash)
			b.FailNow()
		}

		currentBlock.Hash = currentHash
	}
}
