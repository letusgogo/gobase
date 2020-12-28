package util

import (
	"testing"
)

func Test(t *testing.T) {
	aesCBCEncryption := NewAesEncryption("1234567891234567", "1234567891234567")
	DecryptStr, _ := aesCBCEncryption.Encrypt([]byte("1234567891234567"))
	t.Log("aes加密=" + string(DecryptStr))
	EncryptStr, _ := aesCBCEncryption.Decrypt(DecryptStr)
	t.Log("aes解密=" + string(EncryptStr))

	base64Encrypt, _ := base64Encryption{}.Encrypt(DecryptStr)
	t.Log("base64加密=" + string(base64Encrypt))
	Base64Decrypt, _ := base64Encryption{}.Decrypt(base64Encrypt)
	t.Log("base64解密=" + string(Base64Decrypt))
	EncryptStr1, _ := aesCBCEncryption.Decrypt(Base64Decrypt)
	t.Log("aes解密=" + string(EncryptStr1))

}
func Test1(t *testing.T) {
	crypto := NewCrypto()
	crypto.aes.setOption([]byte("1234567891234567"), []byte("1234567891234567"), crypto.mode.CBC, crypto.pad.PKCS5)
	DecryptStr, _ := crypto.aes.Encrypt([]byte("1234567891234567"))
	t.Log("aes加密=" + string(DecryptStr))
	EncryptStr, _ := crypto.aes.Decrypt(DecryptStr)
	t.Log("aes解密=" + string(EncryptStr))

	base64Encrypt, _ := crypto.enc.base64.Encrypt(DecryptStr)
	t.Log("base64加密=" + string(base64Encrypt))
	Base64Decrypt, _ := crypto.enc.base64.Decrypt(base64Encrypt)
	t.Log("base64解密=" + string(Base64Decrypt))
	EncryptStr1, _ := crypto.aes.Decrypt(Base64Decrypt)
	t.Log("aes解密=" + string(EncryptStr1))
}
