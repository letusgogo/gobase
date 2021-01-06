package util

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	crypto := NewCryptoUtil()
	_ = crypto.Aes.SetOption([]byte("1234567891234567"), []byte("1234567891234567"), new(CBC), new(PKCS7))
	DecryptStr, err := crypto.Aes.Encrypt([]byte("1234567891234567"))
	t.Log(err)
	t.Log("aes加密=" + string(DecryptStr))
	EncryptStr, _ := crypto.Aes.Decrypt(DecryptStr)
	t.Log("aes解密=" + string(EncryptStr))

	base64Encrypt, _ := crypto.Base64.Encrypt(DecryptStr)
	t.Log("base64加密=" + string(base64Encrypt))
	Base64Decrypt, _ := crypto.Base64.Decrypt(base64Encrypt)
	t.Log("base64解密=" + string(Base64Decrypt))
	EncryptStr1, _ := crypto.Aes.Decrypt(Base64Decrypt)
	t.Log("aes解密=" + string(EncryptStr1))
}

func Test_Md5(t *testing.T) {
	crypto := NewCryptoUtil()
	_, err := crypto.Base64.Decrypt([]byte("123456"))
	fmt.Println(err)
}
