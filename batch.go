package pbloom

import (
	"math"
	"github.com/esdb/biter"
)

const BatchPutLocationsPerElement = 4

func BatchPut(hashedElement HashedElement,
	slotMask1 biter.Bits, slotMask2 biter.Bits, slotMask3 biter.Bits,
	pbf1 ParallelBloomFilter, pbf2 ParallelBloomFilter, pbf3 ParallelBloomFilter) [4]uint64 {
	combinedHash := uint64(hashedElement[0])

	combinedHash2 := combinedHash & math.MaxUint64
	pbf1[combinedHash2%uint64(len(pbf1))] |= slotMask1
	pbf2[combinedHash2%uint64(len(pbf2))] |= slotMask2
	location1 := (combinedHash2) % uint64(len(pbf3))
	pbf3[location1] |= slotMask3

	combinedHash += uint64(hashedElement[1])
	combinedHash2 = combinedHash & math.MaxUint64
	pbf1[combinedHash2%uint64(len(pbf1))] |= slotMask1
	pbf2[combinedHash2%uint64(len(pbf2))] |= slotMask2
	location2 := combinedHash2 % uint64(len(pbf3))
	pbf3[location2] |= slotMask3

	combinedHash += uint64(hashedElement[1])
	combinedHash2 = combinedHash & math.MaxUint64
	pbf1[combinedHash2%uint64(len(pbf1))] |= slotMask1
	pbf2[combinedHash2%uint64(len(pbf2))] |= slotMask2
	location3 := combinedHash2 % uint64(len(pbf3))
	pbf3[location3] |= slotMask3

	combinedHash += uint64(hashedElement[1])
	combinedHash2 = combinedHash & math.MaxUint64
	pbf1[combinedHash2%uint64(len(pbf1))] |= slotMask1
	pbf2[combinedHash2%uint64(len(pbf2))] |= slotMask2
	location4 := combinedHash2 % uint64(len(pbf3))
	pbf3[location4] |= slotMask3

	return [4]uint64{location1, location2, location3, location4}
}
