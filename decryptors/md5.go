package decryptors 

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(referenceHash, wordFromFile string) bool{
	h := md5.New()
	h.Write([]byte(wordFromFile))
	encrypted := hex.EncodeToString(h.Sum(nil)) 
	return encrypted == referenceHash
}