package pbloom

import (
	"math"
)

// element => hash1,hash2 => k locations
type Hasher func(element []byte) HashedElement

type HashedElement [2]uint64

type BloomElement []uint64

func (hashedElement HashedElement) Hash(locationsCount uint64) BloomElement {
	kHashedElement := make(BloomElement, locationsCount)
	combinedHash := hashedElement[0]
	for i := uint64(0); i < locationsCount; i++ {
		kHashedElement[i] = (combinedHash & math.MaxUint64) % uint64(locationsCount)
		combinedHash += hashedElement[1]
	}
	return kHashedElement
}
