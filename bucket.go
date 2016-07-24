// zheng-ji.info
package cuckoo

import (
//"fmt"
)

// 指纹
type Fp [FingerSize]byte
type Bucket [SlotSize]Fp

var Empty = Fp{0}

func (bk *Bucket) insert(fp Fp) bool {
	for index, vfp := range bk {
		if vfp == Empty {
			bk[index] = fp
			return true
		}
	}
	return false
}

func (bk *Bucket) del(fp Fp) bool {
	for index, vfp := range bk {
		if vfp == fp {
			bk[index] = Empty
			return true
		}
	}
	return false
}

func (bk *Bucket) lookupIndex(fp Fp) int {
	for index, vfp := range bk {
		if vfp == fp {
			return index
		}
	}
	return NotFound
}
