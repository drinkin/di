package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/big"

	"github.com/drinkin/di/lg"
)

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	lg.Check(err)

	return b
}

func Base64(n int) string {
	return base64.URLEncoding.EncodeToString(Bytes(n))
}

func String(n int) string {
	return hex.EncodeToString(Bytes(n))
}

func Int(min, max int) int {
	return int(Int64(min, max))
}

func Int64(min, max int) int64 {
	b, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	lg.Check(err)
	return int64(min) + b.Int64()
}

func SelectString(choices []string) string {
	i := Int64(0, len(choices))
	return choices[i]
}

func SelectInt(choices []int) int {
	i := Int64(0, len(choices))
	return choices[i]
}

func SelectInt64(choices []int64) int64 {
	i := Int64(0, len(choices))
	return choices[i]
}

func SelectFloat64(choices []float64) float64 {
	i := Int64(0, len(choices))
	return choices[i]
}
