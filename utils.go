package alidayu

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

//hmac_md5
func Hmac(s string, key string) string {
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(s))
	return hex.EncodeToString(mac.Sum(nil))
}

//md5
func Md5(key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

func GetInt64Str(d int64) string {
	return strconv.FormatInt(d, 10)
}
