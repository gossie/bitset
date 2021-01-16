package bitset

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	b := BitSet{}
	b.Set(0)
	b.Set(2)
	if !b.IsSet(0) {
		t.Fail()
	}
	if b.IsSet(1) {
		t.Fail()
	}
	if !b.IsSet(2) {
		t.Fail()
	}
}

func TestBoundaries(t *testing.T) {
	b := BitSet{}
	b.Set(0)
	b.Set(7)
	if !b.IsSet(0) {
		t.Fail()
	}
	if b.IsSet(1) {
		t.Fail()
	}
	if b.IsSet(2) {
		t.Fail()
	}
	if b.IsSet(3) {
		t.Fail()
	}
	if b.IsSet(4) {
		t.Fail()
	}
	if b.IsSet(5) {
		t.Fail()
	}
	if b.IsSet(6) {
		t.Fail()
	}
	if !b.IsSet(7) {
		t.Fail()
	}
}

func TestBitsetBiggerThan64Entries(t *testing.T) {
	b := BitSet{}
	b.Set(67)
	b.Set(69)

	for i := uint(0); i < 67; i++ {
		if b.IsSet(i) {
			t.Fail()
		}
	}

	if !b.IsSet(67) {
		t.Fail()
	}
	if b.IsSet(68) {
		t.Fail()
	}
	if !b.IsSet(69) {
		t.Fail()
	}
}

func TestBitsetBiggerThan64EntriesWithCopy(t *testing.T) {
	old := BitSet{}
	old.Set(67)
	old.Set(69)

	b := From(old.Bytes())

	for i := uint(0); i < 67; i++ {
		if b.IsSet(i) {
			t.Fail()
		}
	}

	if !b.IsSet(67) {
		t.Fail()
	}
	if b.IsSet(68) {
		t.Fail()
	}
	if !b.IsSet(69) {
		t.Fail()
	}
}

func TestToString(t *testing.T) {
	bitset := BitSet{}
	bitset.Set(0)
	bitset.Set(2)
	bitset.Set(4)
	bitset.Set(6)
	bitset.Set(8)
	bitset.Set(10)
	bitset.Set(12)
	bitset.Set(14)

	if bitset.String() != "1010101010101010" {
		t.Fatalf("bitset = %v", bitset)
	}
}

func TestAnd(t *testing.T) {
	bitset1 := BitSet{}
	bitset1.Set(0)
	bitset1.Set(2)
	bitset1.Set(4)
	bitset1.Set(6)
	bitset1.Set(8)
	bitset1.Set(10)
	bitset1.Set(12)
	bitset1.Set(14)

	bitset2 := BitSet{}
	bitset2.Set(2)
	bitset2.Set(3)
	bitset2.Set(5)
	bitset2.Set(12)

	if bitset1.And(&bitset2).String() != "0010000000001000" {
		t.Fatalf("%v and %v = %v", bitset1, bitset2, bitset1.And(&bitset2))
	}
}

func TestAndOtherIsLonger(t *testing.T) {
	bitset1 := BitSet{}
	bitset1.Set(0)
	bitset1.Set(2)
	bitset1.Set(4)
	bitset1.Set(6)
	bitset1.Set(8)
	bitset1.Set(10)
	bitset1.Set(12)
	bitset1.Set(14)

	bitset2 := BitSet{}
	bitset2.Set(2)
	bitset2.Set(3)
	bitset2.Set(5)
	bitset2.Set(12)
	bitset2.Set(19)

	if bitset1.And(&bitset2).String() != "0010000000001000" {
		t.Fatalf("%v and %v = %v", bitset1, bitset2, bitset1.And(&bitset2))
	}
}

func TestAndThisIsLonger(t *testing.T) {
	bitset1 := BitSet{}
	bitset1.Set(0)
	bitset1.Set(2)
	bitset1.Set(4)
	bitset1.Set(6)
	bitset1.Set(8)
	bitset1.Set(10)
	bitset1.Set(12)
	bitset1.Set(14)
	bitset1.Set(19)

	bitset2 := BitSet{}
	bitset2.Set(2)
	bitset2.Set(3)
	bitset2.Set(5)
	bitset2.Set(12)

	if bitset1.And(&bitset2).String() != "0010000000001000" {
		t.Fatalf("%v and %v = %v", bitset1, bitset2, bitset1.And(&bitset2))
	}
}
func TestOr(t *testing.T) {
	bitset1 := BitSet{}
	bitset1.Set(0)
	bitset1.Set(2)
	bitset1.Set(4)
	bitset1.Set(6)
	bitset1.Set(8)
	bitset1.Set(10)
	bitset1.Set(12)
	bitset1.Set(14)

	bitset2 := BitSet{}
	bitset2.Set(2)
	bitset2.Set(3)
	bitset2.Set(5)
	bitset2.Set(12)

	if bitset1.Or(&bitset2).String() != "1011111010101010" {
		t.Fatalf("%v or %v = %v", bitset1, bitset2, bitset1.Or(&bitset2))
	}
}

func TestOrOtherIsLonger(t *testing.T) {
	bitset1 := BitSet{}
	bitset1.Set(0)
	bitset1.Set(2)
	bitset1.Set(4)
	bitset1.Set(6)
	bitset1.Set(8)
	bitset1.Set(10)
	bitset1.Set(12)
	bitset1.Set(14)

	bitset2 := BitSet{}
	bitset2.Set(2)
	bitset2.Set(3)
	bitset2.Set(5)
	bitset2.Set(12)
	bitset2.Set(19)

	if bitset1.Or(&bitset2).String() != "101111101010101000010000" {
		t.Fatalf("%v or %v = %v", bitset1, bitset2, bitset1.Or(&bitset2))
	}
}

func TestOrThisIsLonger(t *testing.T) {
	bitset1 := BitSet{}
	bitset1.Set(0)
	bitset1.Set(2)
	bitset1.Set(4)
	bitset1.Set(6)
	bitset1.Set(8)
	bitset1.Set(10)
	bitset1.Set(12)
	bitset1.Set(14)
	bitset1.Set(19)

	bitset2 := BitSet{}
	bitset2.Set(2)
	bitset2.Set(3)
	bitset2.Set(5)
	bitset2.Set(12)

	if bitset1.Or(&bitset2).String() != "101111101010101000010000" {
		t.Fatalf("%v or %v = %v", bitset1, bitset2, bitset1.Or(&bitset2))
	}
}
