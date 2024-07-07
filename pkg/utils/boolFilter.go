package utils

import "hash/fnv"

func BoolHash32(data []byte) uint32 {
	var res uint32
	hash := fnv.New32()
	hash.Write(data)
	res = hash.Sum32()

	return res
}
