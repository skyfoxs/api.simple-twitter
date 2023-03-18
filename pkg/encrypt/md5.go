package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5str(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
