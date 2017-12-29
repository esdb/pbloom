package pbloom

import (
	"github.com/esdb/biter"
	"math"
)

const SlotNotFound = biter.NotFound

// 64 slot bloom filter
type ParallelBloomFilter struct {
	locationsPerElement int
	locations           []biter.Bits
}

func New(locationsCount, locationsPerElement int) *ParallelBloomFilter {
	return &ParallelBloomFilter{
		locationsPerElement: locationsPerElement,
		locations:           make([]biter.Bits, locationsCount),
	}
}

func (pbf *ParallelBloomFilter) Put(hasher Hasher, slot biter.Bits, element []byte) {
	hashedElement := hasher(element)
	combinedHash := hashedElement[0]
	locationsCount := uint64(len(pbf.locations))
	for i := 0; i < pbf.locationsPerElement; i++ {
		pbf.locations[(combinedHash&math.MaxUint64)%locationsCount] |= slot
		combinedHash += hashedElement[1]
	}
}

func (pbf *ParallelBloomFilter) Find(hasher Hasher, element []byte) biter.Bits {
	hashedElement := hasher(element)
	combinedHash := hashedElement[0]
	locationsCount := uint64(len(pbf.locations))
	result := biter.SetAllBits
	for i := 0; i < pbf.locationsPerElement; i++ {
		result &= pbf.locations[(combinedHash&math.MaxUint64)%locationsCount]
		combinedHash += hashedElement[1]
	}
	return result
}
