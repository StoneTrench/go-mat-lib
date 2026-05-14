package lib

import (
	"crypto/rand"
	"encoding/binary"
)

type UUID [4]uint32

func NewUUID() (UUID, error) {
	var b [16]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return UUID{}, err_append("UUID failed to initialize random", err)
	}

	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	var uuid UUID
	uuid[0] = binary.BigEndian.Uint32(b[0:4])
	uuid[1] = binary.BigEndian.Uint32(b[4:8])
	uuid[2] = binary.BigEndian.Uint32(b[8:12])
	uuid[3] = binary.BigEndian.Uint32(b[12:16])

	return uuid, nil
}
func (s UUID) ToFixedSlice() [4]int32 {
	return [4]int32{
		int32(s[0]),
		int32(s[1]),
		int32(s[2]),
		int32(s[3]),
	}
}
