package utils

import "strconv"

func StringToUint(s string) uint64 {
	i, _ := strconv.Atoi(s)
	return uint64(i)
}