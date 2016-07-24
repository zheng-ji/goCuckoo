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
		buckets[i] = [SlotSize]Fp{}
	}
	return &CuckooFilter{
		buckets: buckets,
		num:     0,
	}
}

func (filter *CuckooFilter) Find(data []byte) bool {
	fp := genFp(data)
	firstIndex := genFirstIndex(fp, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(fp, uint(len(filter.buckets)))

	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]

	if bk1.lookupIndex(fp) != NotFound || bk2.lookupIndex(fp) != NotFound {
		return true
	}
	return false
}

func (filter *CuckooFilter) Insert(data []byte) bool {
	fp := genFp(data)
	firstIndex := genFirstIndex(fp, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(fp, uint(len(filter.buckets)))
	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]
	if bk1.insert(fp) || bk2.insert(fp) {
		filter.num++
		return true
	}
	return filter.reinsert(fp, backupIndex)
}

func (filter *CuckooFilter) reinsert(fp Fp, i uint) bool {
	for k := 0; k < MaxCuckooCount; k++ {
		j := rand.Intn(SlotSize)
		tmpfp := fp
		fp = filter.buckets[i][j]
		filter.buckets[i][j] = tmpfp
		i = genBackupIndex(fp, uint(len(filter.buckets)))
		bk := filter.buckets[i]
		if bk.insert(fp) {
			filter.num++
			return true
		}
	}
	return false
}

func (filter *CuckooFilter) Del(data []byte) bool {
	fp := genFp(data)
	firstIndex := genFirstIndex(fp, uint(len(filter.buckets)))
	backupIndex := genBackupIndex(fp, uint(len(filter.buckets)))
	bk1 := &filter.buckets[firstIndex]
	bk2 := &filter.buckets[backupIndex]
	return bk1.del(fp) || bk2.del(fp)
}

func (filter *CuckooFilter) Count() int {
	return filter.num
}
