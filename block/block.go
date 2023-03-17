package block

import "time"

type Block struct {
	Data      BlockData
	Hash      string
	PrevHash  string
	Timestamp time.Time
	Pow       int
}
