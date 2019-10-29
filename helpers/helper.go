package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// SaltMaker 生成盐
func SaltMaker(saltLen int) string {
	str := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)

	if saltLen < 5 {
		saltLen = 5
	} else if saltLen > 32 {
		saltLen = 32
	}
	result := make([]byte, saltLen)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < saltLen; i++ {
		index := rand.Intn(len(bytes))
		result[i] = bytes[index]
	}
	return string(result)
}

//MD5 MD5
func MD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
