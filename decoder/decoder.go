package decoder

import (
	"hash/crc32"
)

func DecodeURL(x string) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	result := crc32.Checksum([]byte(x), table)
	return result
}