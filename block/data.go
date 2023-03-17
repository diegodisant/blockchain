package block

type BlockData = map[string]interface{}

func BuildBlockData(from, to string, amount float64) BlockData {
	return BlockData{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
}
