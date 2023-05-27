package functions

import (
	"dehasher/decryptors"
)

func SelectAlgorithm(referenceHash, wordFromFile, algo string) bool {
	var returnValue bool
	switch algo {
	case "SHA256", "Sha256", "sha256":
		returnValue = decryptors.Sha256(referenceHash, wordFromFile)
	case "MD5", "Md5", "md5":
		returnValue = decryptors.Md5(referenceHash, wordFromFile)
	}
	return returnValue
}