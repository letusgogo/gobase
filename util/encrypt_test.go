package util

import (
	"testing"
)

func Test(t *testing.T) {
	aesCBCEncryption := NewAesCBCEncryption("1234567891234567", "1234567891234567")
	DecryptStr, _ := aesCBCEncryption.Encrypt("1234567891234567")
	t.Log(DecryptStr)
	EncryptStr, _ := aesCBCEncryption.Decrypt(DecryptStr)

	t.Log(EncryptStr)

}
