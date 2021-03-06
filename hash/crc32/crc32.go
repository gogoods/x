package crc32

import (
	"hash/crc32"
)

func EncryptString(v string) uint32 {

	//return crc32.Checksum([]byte(v), crc32.MakeTable(crc32.IEEE))
	return crc32.ChecksumIEEE([]byte(v))
}

func EncryptBytes(v []byte) uint32 {
	return crc32.ChecksumIEEE(v)
}
