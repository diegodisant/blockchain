package block

import "fmt"

type BlockEncodeError struct {
	CurrentBlock *Block
	EncodeError  error
}

func NewBlockEncodeError(block *Block, encodeError error) *BlockEncodeError {
	return &BlockEncodeError{
		CurrentBlock: block,
		EncodeError:  encodeError,
	}
}

var _ error = (*BlockEncodeError)(nil)

func (err *BlockEncodeError) Error() string {
	return fmt.Sprintf(
		"Unable to decode block data with cause: %s\n with value: %s\n",
		err.EncodeError.Error(),
		err.CurrentBlock.Data,
	)
}
