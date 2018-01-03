package pbloom

import (
	"testing"
	"github.com/esdb/biter"
	"github.com/stretchr/testify/require"
)

func Test_one_location(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv, 256, 1)
	pbf := strategy.New()
	slot17 := biter.SetBits[17]
	strategy.Put(pbf, slot17, []byte("hello"))
	result := strategy.Find(pbf, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	should.Equal(biter.Slot(17), result.ScanForward()())
	result = strategy.Find(pbf, []byte("world2"))
	should.Equal(biter.Bits(0), result)
}

func Test_seven_locations(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv,256, 7)
	pbf := strategy.New()
	slot17 := biter.SetBits[17]
	strategy.Put(pbf, slot17, []byte("hello"))
	result := strategy.Find(pbf, []byte("hello"))
	should.NotEqual(0, result)
	should.Equal(biter.Slot(17), result.ScanForward()())
}

func Test_two_slots(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv, 256, 1)
	pbf := strategy.New()
	slot17 := biter.SetBits[17]
	strategy.Put(pbf, slot17, []byte("hello"))
	slot18 := biter.SetBits[18]
	strategy.Put(pbf, slot18, []byte("hello"))
	result := strategy.Find(pbf, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	iter := result.ScanForward()
	should.Equal(biter.Slot(17), iter())
	should.Equal(biter.Slot(18), iter())
	should.Equal(SlotNotFound, iter())
}

func Test_pre_hashing(t *testing.T) {
	should := require.New(t)
	strategy := NewHashingStrategy(HasherFnv, 256, 1)
	pbf := strategy.New()
	element := strategy.Hash([]byte("hello"))
	slot18 := biter.SetBits[18]
	pbf.Put(slot18, element)
	result := pbf.Find(element)
	should.NotEqual(biter.Bits(0), result)
	iter := result.ScanForward()
	should.Equal(biter.Slot(18), iter())
	should.Equal(biter.Slot(64), iter())
}

func Test_batch_put(t *testing.T) {
	should := require.New(t)
	strategy1 := NewHashingStrategy(HasherFnv, 256, BatchPutLocationsPerElement)
	strategy2 := NewHashingStrategy(HasherFnv, 1024, BatchPutLocationsPerElement)
	strategy3 := NewHashingStrategy(HasherFnv, 7777, BatchPutLocationsPerElement)
	pbf1 := strategy1.New()
	pbf2 := strategy2.New()
	pbf3 := strategy3.New()
	hashedElement := HasherFnv([]byte("hello"))
	slot18 := biter.SetBits[18]
	BatchPut(hashedElement, slot18, slot18, slot18, pbf1, pbf2, pbf3)
	result := strategy2.Find(pbf2, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	iter := result.ScanForward()
	should.Equal(biter.Slot(18), iter())
	should.Equal(biter.Slot(64), iter())
	result = strategy3.Find(pbf3, []byte("hello"))
	should.NotEqual(biter.Bits(0), result)
	iter = result.ScanForward()
	should.Equal(biter.Slot(18), iter())
	should.Equal(biter.Slot(64), iter())
}
