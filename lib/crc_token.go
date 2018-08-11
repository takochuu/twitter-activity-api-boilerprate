package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func CreateCRCToken(crcToken string, cs string) string {
	mac := hmac.New(sha256.New, []byte(cs))
	mac.Write([]byte(crcToken))
	return "sha256=" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
