package pbloom

import (
	"math"
)

// element => hash1,hash2 => k locations
type Hasher func(element []byte) HashedElement

type HashedElement [2]uint64

type KHashedElement []uint64

func (hashedElement HashedElement) HashToKLocations(locationsCount int) KHashedElement {
	kHashedElement := make(KHashedElement, locationsCount)
	combinedHash := hashedElement[0]
	for i := 0; i < locationsCount; i++ {
		kHashedElement[i] = (combinedHash & math.MaxUint64) % uint64(locationsCount)
		combinedHash += hashedElement[1]
	}
	return kHashedElement
}
