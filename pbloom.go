package pbloom

import (
	"github.com/esdb/biter"
	"math"
)

const SlotNotFound = biter.NotFound

// 64 slot bloom filter
type ParallelBloomFilter []biter.Bits

type HashingStrategy struct {
	hasher              Hasher
	locationsPerElement uint64
	locationsCount      uint64
}

func NewHashingStrategy(hasher Hasher, locationsCount, locationsPerElement uint64) *HashingStrategy {
	return &HashingStrategy{
		hasher:              hasher,
		locationsPerElement: locationsPerElement,
		locationsCount:      locationsCount,
	}
}

func (strategy *HashingStrategy) HashStage1(element []byte) HashedElement {
	return strategy.hasher(element)
}

func (strategy *HashingStrategy) HashStage2(hashedElement HashedElement) BloomElement {
	return hashedElement.Hash(strategy.locationsPerElement)
}

func (strategy *HashingStrategy) Hash(element []byte) BloomElement {
	return strategy.hasher(element).Hash(strategy.locationsPerElement)
}

func (strategy *HashingStrategy) New() ParallelBloomFilter {
	return make(ParallelBloomFilter, strategy.locationsCount)
}

func (strategy *HashingStrategy) Put(pbf ParallelBloomFilter, slot biter.Bits, element []byte) {
	hashedElement := strategy.hasher(element)
	combinedHash := hashedElement[0]
	locationsCount := strategy.locationsCount
	for i := uint64(0); i < strategy.locationsPerElement; i++ {
		pbf[(combinedHash&math.MaxUint64)%locationsCount] |= slot
		combinedHash += hashedElement[1]
	}
}

func (strategy *HashingStrategy) Find(pbf ParallelBloomFilter, element []byte) biter.Bits {
	hashedElement := strategy.hasher(element)
	combinedHash := hashedElement[0]
	locationsCount := strategy.locationsCount
	result := biter.SetAllBits
	for i := uint64(0); i < strategy.locationsPerElement; i++ {
		result &= pbf[(combinedHash&math.MaxUint64)%locationsCount]
		combinedHash += hashedElement[1]
	}
	return result
}

func (pbf ParallelBloomFilter) Put(slot biter.Bits, element BloomElement) {
	for _, location := range element {
		pbf[location] |= slot
	}
}

func (pbf ParallelBloomFilter) Find(element BloomElement) biter.Bits {
	result := biter.SetAllBits
	for _, location := range element {
		result &= pbf[location]
	}
	return result
}
