package storage



func GetKeyHash(key string) int {
	hash := 0
	l := len(key)
	for i := 0; i < l; i++ {
		hash += int(key[i])
	}
	return l + hash
}
