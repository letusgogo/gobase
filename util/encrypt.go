package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

///////////////////////////加密解密////////////////////////////////
type Encryption interface {
	Encrypt(encodeStr string) (string, error)
	Decrypt(decodeStr string) (string, error)
}

//=======================Aes CBC实现===========================
type AesCBCEncryption struct {
	SecretKey string
	SecretIv  string
}

func NewAesCBCEncryption(secretKey string, secretIv string) *AesCBCEncryption {
	return &AesCBCEncryption{SecretKey: secretKey, SecretIv: secretIv}
}

func (aesCBC *AesCBCEncryption) Encrypt(encodeStr string) (string, error) {

	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher([]byte(aesCBC.SecretKey))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(aesCBC.SecretIv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func (aesCBC *AesCBCEncryption) Decrypt(decodeStr string) (string, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(aesCBC.SecretKey))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(aesCBC.SecretIv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//========================== md5 ===============================
type Md5EncryptionImpl struct {
}

func (md5Impl Md5EncryptionImpl) Encrypt(encodeStr string) (string, error) {
	h := md5.New()
	h.Write([]byte(encodeStr))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (md5Impl Md5EncryptionImpl) Decrypt(decodeStr string) (string, error) {
	panic("please implement me")
}
