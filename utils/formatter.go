package utils

import "strconv"

func StringToUint(s string) uint64 {
	i, _ := strconv.Atoi(s)
	return uint64(i)
}

func StringToUint32(s string) uint32 {
	i, _ := strconv.ParseUint(s, 10, 32)
	return uint32(i)
}