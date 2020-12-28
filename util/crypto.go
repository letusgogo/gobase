package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
)

///////////////////////////加密的入口 ////////////////////////////////
type CryptoUtil struct {
	Aes    AesEncryption
	Base64 Base64Encryption
	Md5    Md5Encryption
}

func NewCryptoUtil() *CryptoUtil {
	crypto := new(CryptoUtil)
	return crypto
}

///////////////////////////////补全方式的接口和实现////////////////////////
type Pad interface {
	Padding(ciphertext []byte, blockSize int) []byte
	UnPadding(origData []byte) []byte
}

type PKCS7 struct {
}

func (P PKCS7) Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func (P PKCS7) UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

/////////////////////////////加密方式的接口和实现///////////////////////
type Mode interface {
	Encrypt(iv []byte, block cipher.Block, encodeBytes []byte) ([]byte, error)
	Decrypt(iv []byte, block cipher.Block, decodeBytes []byte) ([]byte, error)
}

type CBC struct {
}

func (cbc *CBC) Encrypt(iv []byte, block cipher.Block, encodeBytes []byte) ([]byte, error) {
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)
	return crypted, nil
}

func (cbc *CBC) Decrypt(iv []byte, block cipher.Block, decodeBytes []byte) ([]byte, error) {
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	return origData, nil
}

//////////////////////////加密的接口和实现类//////////////////////////////
type Encryption interface {
	Encrypt(encodeBytes []byte) ([]byte, error)
	Decrypt(decodeBytes []byte) ([]byte, error)
}

//=======================Aes实现===========================
type AesEncryption struct {
	//key
	secretKey []byte
	//iv
	secretIv []byte
	//加密方式
	mode Mode
	//补全方式
	pad Pad
}

func (aesEncryption *AesEncryption) JudgeOption() error {
	SecretKey := aesEncryption.secretKey
	SecretIv := aesEncryption.secretIv
	if len(SecretKey) != 16 && len(SecretKey) != 24 && len(SecretKey) != 32 {
		return errors.New("key length must be 16 or 24 or 32 byte")
	}
	if len(SecretKey) != len(SecretIv) {
		return errors.New("key and iv must be equal")
	}
	if aesEncryption.mode == nil || aesEncryption.pad == nil {
		return errors.New("Mode and Pad not empty")
	}
	return nil
}

func (aesEncryption *AesEncryption) SetOption(SecretKey []byte, SecretIv []byte, mode Mode, pad Pad) error {
	aesEncryption.secretKey = SecretKey
	aesEncryption.secretIv = SecretIv
	aesEncryption.mode = mode
	aesEncryption.pad = pad
	return aesEncryption.JudgeOption()
}

func (aesEncryption *AesEncryption) Encrypt(encodeBytes []byte) ([]byte, error) {
	if err := aesEncryption.JudgeOption(); err != nil {
		return nil, err
	}
	//根据key 生成密文
	block, err := aes.NewCipher(aesEncryption.secretKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encodeBytes = aesEncryption.pad.Padding(encodeBytes, blockSize)
	crypted, _ := aesEncryption.mode.Encrypt(aesEncryption.secretIv, block, encodeBytes)
	return crypted, nil
}

func (aesEncryption *AesEncryption) Decrypt(decodeBytes []byte) ([]byte, error) {
	if err := aesEncryption.JudgeOption(); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesEncryption.secretKey)
	if err != nil {
		return nil, err
	}
	origData, _ := aesEncryption.mode.Decrypt(aesEncryption.secretIv, block, decodeBytes)
	origData = aesEncryption.pad.UnPadding(origData)
	return origData, nil
}

//========================== Md5 ===============================
type Md5Encryption struct {
}

func (md5Encryption Md5Encryption) Encrypt(encodeBytes []byte) ([]byte, error) {
	h := md5.New()
	h.Write(encodeBytes)
	return h.Sum(nil), nil
	//return hex.EncodeToString(h.Sum(nil)), nil
}

func (md5Encryption Md5Encryption) Decrypt(decodeBytes []byte) ([]byte, error) {
	panic("please implement me")
}

//============================Base64===========================
type Base64Encryption struct {
}

func (base64Encryption Base64Encryption) Encrypt(encodeBytes []byte) ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(encodeBytes)), nil
}

func (base64Encryption Base64Encryption) Decrypt(decodeBytes []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(string(decodeBytes))
	if err != nil {
		return nil, err
	}
	return decodeBytes, nil
}
