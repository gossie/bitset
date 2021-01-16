package bitset

import "strings"

// BitSet is a bitset.
type BitSet []byte

// From creates a new bitset from the given slice.
func From(data []byte) BitSet {
	return data
}

func (b BitSet) String() string {
	stringBuilder := strings.Builder{}
	for _, bytes := range b {
		mask := byte(1)
		for i := 0; i < 8; i++ {
			if (bytes & (mask << i)) != 0 {
				stringBuilder.WriteString("1")
			} else {
				stringBuilder.WriteString("0")
			}
		}
	}
	return stringBuilder.String()
}

// IsSet returns true if the bit at the given index is set.
func (b *BitSet) IsSet(index uint) bool {
	dataIndex := index >> 3
	if dataIndex >= uint(len(*b)) {
		return false
	}
	value := 1 << (index & 7)
	return ((*b)[dataIndex] & byte(value)) > 0
}

// Set sets the bit at the given index.
func (b *BitSet) Set(index uint) {
	dataIndex := index >> 3
	for i := uint(len(*b)); i <= dataIndex; i++ {
		*b = append(*b, 0)
	}
	value := 1 << (index & 7)
	(*b)[dataIndex] |= byte(value)
}

// Bytes returns the a slice of byte.
func (b *BitSet) Bytes() []byte {
	return *b
}
