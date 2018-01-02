package pbloom

import (
	"math"
	"github.com/esdb/biter"
)

func BatchPut2(slotMask biter.Bits, hashedElement HashedElement,
	locationsPerElement uint64,
	pbf1 ParallelBloomFilter, pbf2 ParallelBloomFilter) {
	combinedHash := uint64(hashedElement[0])
	for i := uint64(0); i < locationsPerElement; i++ {
		pbf1[(combinedHash&math.MaxUint64)%uint64(len(pbf1))] |= slotMask
		pbf2[(combinedHash&math.MaxUint64)%uint64(len(pbf2))] |= slotMask
		combinedHash += uint64(hashedElement[1])
	}
}

func BatchPut3(slotMask biter.Bits, hashedElement HashedElement,
	locationsPerElement uint64,
	pbf1 ParallelBloomFilter, pbf2 ParallelBloomFilter, pbf3 ParallelBloomFilter) {
	combinedHash := uint64(hashedElement[0])
	for i := uint64(0); i < locationsPerElement; i++ {
		pbf1[(combinedHash&math.MaxUint64)%uint64(len(pbf1))] |= slotMask
		pbf2[(combinedHash&math.MaxUint64)%uint64(len(pbf2))] |= slotMask
		pbf3[(combinedHash&math.MaxUint64)%uint64(len(pbf3))] |= slotMask
		combinedHash += uint64(hashedElement[1])
	}
}
