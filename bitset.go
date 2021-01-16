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

func (b *BitSet) len() int {
	return len(*b)
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

// And performs an and between two bitsets.
func (b *BitSet) And(other *BitSet) BitSet {
	result := make([]byte, 0)
	shorter, longer := determineShorterAndLongerBitset(b, other)
	for index, bits := range *shorter {
		result = append(result, bits&(*longer)[index])
	}
	return result
}

// Or performs an and between two bitsets.
func (b *BitSet) Or(other *BitSet) BitSet {
	result := make([]byte, 0)
	shorter, longer := determineShorterAndLongerBitset(b, other)
	for index, bits := range *shorter {
		result = append(result, bits|(*longer)[index])
	}
	return append(result, (*longer)[shorter.len():]...)
}

// Xor performs an and between two bitsets.
func (b *BitSet) Xor(other *BitSet) BitSet {
	result := make([]byte, 0)
	shorter, longer := determineShorterAndLongerBitset(b, other)
	for index, bits := range *shorter {
		result = append(result, bits^(*longer)[index])
	}
	return append(result, (*longer)[shorter.len():]...)
}

func determineShorterAndLongerBitset(b1 *BitSet, b2 *BitSet) (*BitSet, *BitSet) {
	if b1.len() < b2.len() {
		return b1, b2
	}
	return b2, b1
}
