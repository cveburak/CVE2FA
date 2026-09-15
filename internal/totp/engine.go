package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

const (
	Period    = 30
	CodeLen   = 6
	EpochZero = 0
)

func GenerateCode(secret []byte, counter uint64) string {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, counter)

	mac := hmac.New(sha1.New, secret)
	mac.Write(buf)
	hash := mac.Sum(nil)

	offset := hash[len(hash)-1] & 0x0f
	truncated := binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7fffffff

	code := truncated % uint32(math.Pow10(CodeLen))
	return fmt.Sprintf("%06d", code)
}

func GetCounter(t time.Time) uint64 {
	return uint64(t.Unix()) / Period
}

func GetRemainingSeconds(t time.Time) int {
	return Period - (int(t.Unix()) % Period)
}

func GenerateCodeFromTime(secret []byte, t time.Time) string {
	return GenerateCode(secret, GetCounter(t))
}

func ValidateCode(secret []byte, code string, t time.Time) bool {
	counter := GetCounter(t)
	for i := -1; i <= 1; i++ {
		if GenerateCode(secret, counter+uint64(i)) == code {
			return true
		}
	}
	return false
}
