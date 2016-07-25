// zheng-ji.info
package cuckoo

import (
	//"fmt"
	"math/rand"
)

type CuckooFilter struct {
	num     int
	buckets []Bucket
}

func NewCuckooFilter(capacity uint) *CuckooFilter {
	capacity = getCeilingCap(uint64(capacity)) / SlotSize
	if capacity == 0 {
		capacity = 1
	}
	buckets := make([]Bucket, capacity, capacity)
	for i := range buckets {
		buckets[i] = [SlotSize]Signature{}
	}
	return &CuckooFilter{
		buckets: buckets,
		num:     0,
	}
}

func (filter *CuckooFilter) Find(data []byte) bool {
	sign := genSignature(data)
	firstIndex := genFirstIndex(sign, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(sign, uint(len(filter.buckets)))

	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]

	if bk1.lookupIndex(sign) != NotFound || bk2.lookupIndex(sign) != NotFound {
		return true
	}
	return false
}

func (filter *CuckooFilter) Insert(data []byte) bool {
	sign := genSignature(data)
	firstIndex := genFirstIndex(sign, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(sign, uint(len(filter.buckets)))
	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]
	if bk1.insert(sign) || bk2.insert(sign) {
		filter.num++
		return true
	}
	return filter.resolveCollision(sign, backupIndex)
}

func (filter *CuckooFilter) resolveCollision(sign Signature, index uint) bool {
	for i := 0; i < MaxCuckooCount; i++ {
		j := rand.Intn(SlotSize)
		tmpsign := sign
		sign = filter.buckets[index][j]
		filter.buckets[index][j] = tmpsign
		index = genBackupIndex(sign, uint(len(filter.buckets)))
		bk := &filter.buckets[index]
		if bk.insert(sign) {
			filter.num++
			return true
		}
	}
	return false
}

func (filter *CuckooFilter) Del(data []byte) bool {
	sign := genSignature(data)
	firstIndex := genFirstIndex(sign, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(sign, uint(len(filter.buckets)))
	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]
	return bk1.del(sign) || bk2.del(sign)
}

func (filter *CuckooFilter) Size() int {
	return filter.num
}
