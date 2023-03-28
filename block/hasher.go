package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/diegodisant/blockchain/common"
)

var _ common.Hashable = (*Block)(nil)

func (block *Block) CalculateHash() string {
	encodedData, err := json.Marshal(block.Data)

	if err != nil {
		panic(NewBlockEncodeError(
			block,
			err,
		))
	}

	blockHashSalt := block.PrevHash + string(encodedData) + block.Timestamp.String() + strconv.Itoa(block.Pow)
	blockHash := sha256.Sum256([]byte(blockHashSalt))

	return hex.EncodeToString(blockHash[:])
}
