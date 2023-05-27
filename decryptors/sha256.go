package decryptors 

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(referenceHash, wordFromFile string) bool{
	h := sha256.New()
	h.Write([]byte(wordFromFile))
	encrypted := hex.EncodeToString(h.Sum(nil)) 
	return encrypted == referenceHash
}