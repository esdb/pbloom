package pbloom

// fnv is fast for small data
// xxhash is fast for larger data

const (
	offset128Lower  = 0x62b821756295c58d
	offset128Higher = 0x6c62272e07bb0142
	prime128Lower   = 0x13b
	prime128Shift   = 24
)

// copied stdlib hash/fnv
func HasherFnv(element []byte) (s HashedElement) {
	s[0] = offset128Higher
	s[1] = offset128Lower
	for _, c := range element {
		s[1] ^= uint64(c)
		// Compute the multiplication in 4 parts to simplify carrying
		s1l := (s[1] & 0xffffffff) * prime128Lower
		s1h := (s[1] >> 32) * prime128Lower
		s0l := (s[0]&0xffffffff)*prime128Lower + (s[1]&0xffffffff)<<prime128Shift
		s0h := (s[0]>>32)*prime128Lower + (s[1]>>32)<<prime128Shift
		// Carries
		s1h += s1l >> 32
		s0l += s1h >> 32
		s0h += s0l >> 32
		// Update the values
		s[1] = (s1l & 0xffffffff) + (s1h << 32)
		s[0] = (s0l & 0xffffffff) + (s0h << 32)
	}
	return s
}
