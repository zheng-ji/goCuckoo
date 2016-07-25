// zheng-ji.info
package cuckoo

type Signature [SignatureSize]byte
type Bucket [SlotSize]Signature

var Empty = Signature{0}

func (bk *Bucket) insert(sign Signature) bool {
	for index, vsign := range bk {
		if vsign == Empty {
			bk[index] = sign
			return true
		}
	}
	return false
}

func (bk *Bucket) del(sign Signature) bool {
	for index, vsign := range bk {
		if vsign == sign {
			bk[index] = Empty
			return true
		}
	}
	return false
}

func (bk *Bucket) lookupIndex(sign Signature) int {
	for index, vsign := range bk {
		if vsign == sign {
			return index
		}
	}
	return NotFound
}
