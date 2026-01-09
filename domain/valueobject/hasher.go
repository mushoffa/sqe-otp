package valueobject

import (
	"crypto/sha256"
	"encoding/base64"
)

func Hasher(payload string) string {
	hasher := sha256.New()
	hasher.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
