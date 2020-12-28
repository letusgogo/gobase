package util

import (
	"testing"
)

func Test(t *testing.T) {
	crypto := NewCryptoUtil()
	crypto.aes.SetOption([]byte("1234567891234567"), []byte("1234567891234567"), new(CBC), new(PKCS5))
	DecryptStr, err := crypto.aes.Encrypt([]byte("1234567891234567"))
	t.Log(err)
	t.Log("aes加密=" + string(DecryptStr))
	EncryptStr, _ := crypto.aes.Decrypt(DecryptStr)
	t.Log("aes解密=" + string(EncryptStr))

	base64Encrypt, _ := crypto.base64.Encrypt(DecryptStr)
	t.Log("base64加密=" + string(base64Encrypt))
	Base64Decrypt, _ := crypto.base64.Decrypt(base64Encrypt)
	t.Log("base64解密=" + string(Base64Decrypt))
	EncryptStr1, _ := crypto.aes.Decrypt(Base64Decrypt)
	t.Log("aes解密=" + string(EncryptStr1))
}
