package encode

import (

	"time"
    
    "encoding/base64"

	"crypto/sha256"

	"crypto/hmac"
    
    "strconv"
)

func unixDate() int64 {

	// Convert time now into a unix int using base 10
	// the result would be something like: 1424862288
    // strconv.FormatInt(time.Now().Unix(), 10)
	unix := time.Now().Unix()

    return unix + 100
}

func HmacSha256Encode(key string, secret string) (string, int64) {
    
    e := unixDate()

    // Convert e (int64) into a string using base 10
    message := key + "\n" + strconv.FormatInt(e, 10)
    
	bits, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {

		return "Toasted", 0
	}

	h := hmac.New(sha256.New, bits)

	h.Write([]byte(message))

	return base64.URLEncoding.EncodeToString(h.Sum(nil)), e   
}