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
type crypto struct {
	mode modeMap
	pad  padMap
	aes  aesEncryption
	//encoding的意思
	enc encMap
}

func NewCrypto() *crypto {
	crypto := new(crypto)
	crypto.mode.CBC = &cBC{}
	crypto.pad.PKCS5 = &pKCS5{}
	return crypto
}

type encMap struct {
	base64 base64Encryption
	md5    md5Encryption
}

type padMap struct {
	PKCS5 pad
}

type modeMap struct {
	CBC mode
}

///////////////////////////////补全方式的接口和实现////////////////////////
type pad interface {
	Padding(ciphertext []byte, blockSize int) []byte
	UnPadding(origData []byte) []byte
}

type pKCS5 struct {
}

func (P pKCS5) Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func (P pKCS5) UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

/////////////////////////////加密方式的接口和实现///////////////////////
type mode interface {
	Encrypt(iv []byte, block cipher.Block, encodeStr []byte) ([]byte, error)
	Decrypt(iv []byte, block cipher.Block, encodeStr []byte) ([]byte, error)
}

type cBC struct {
}

func (cbc *cBC) Encrypt(iv []byte, block cipher.Block, encodeStr []byte) ([]byte, error) {
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeStr))
	blockMode.CryptBlocks(crypted, encodeStr)
	return crypted, nil
}

func (cbc *cBC) Decrypt(iv []byte, block cipher.Block, decodeStr []byte) ([]byte, error) {
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeStr))
	blockMode.CryptBlocks(origData, decodeStr)
	return origData, nil
}

//////////////////////////加密的接口和实现类//////////////////////////////
type Encryption interface {
	Encrypt(encodeStr []byte) ([]byte, error)
	Decrypt(decodeStr []byte) ([]byte, error)
}

//=======================Aes实现===========================
type aesEncryption struct {
	//key
	secretKey string
	//iv
	secretIv string
	//加密方式
	mode mode
	//补全方式
	pad pad
}

func NewAesEncryption(secretKey string, secretIv string) *aesEncryption {
	return &aesEncryption{secretKey: secretKey, secretIv: secretIv}
}

func (aesEncryption *aesEncryption) setOption(SecretKey []byte, SecretIv []byte, mode mode, pad pad) error {
	if len(SecretKey) != 16 && len(SecretKey) != 24 && len(SecretKey) != 32 {
		return errors.New("")
	}
	if len(SecretKey) != len(SecretIv) {
		return errors.New("")
	}
	aesEncryption.secretKey = string(SecretKey)
	aesEncryption.secretIv = string(SecretIv)
	aesEncryption.mode = mode
	aesEncryption.pad = pad
	return nil
}

func (aesEncryption *aesEncryption) Encrypt(encodeStr []byte) ([]byte, error) {
	//根据key 生成密文
	block, err := aes.NewCipher([]byte(aesEncryption.secretKey))
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encodeStr = aesEncryption.pad.Padding(encodeStr, blockSize)
	crypted, _ := aesEncryption.mode.Encrypt([]byte(aesEncryption.secretIv), block, encodeStr)
	return crypted, nil
}

func (aesEncryption *aesEncryption) Decrypt(decodeStr []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(aesEncryption.secretKey))
	if err != nil {
		return nil, err
	}
	origData, _ := aesEncryption.mode.Decrypt([]byte(aesEncryption.secretIv), block, decodeStr)
	origData = aesEncryption.pad.UnPadding(origData)
	return origData, nil
}

//========================== md5 ===============================
type md5Encryption struct {
}

func (md5Encryption md5Encryption) Encrypt(encodeStr []byte) ([]byte, error) {
	h := md5.New()
	h.Write(encodeStr)
	return h.Sum(nil), nil
	//return hex.EncodeToString(h.Sum(nil)), nil
}

func (md5Encryption md5Encryption) Decrypt(decodeStr []byte) ([]byte, error) {
	panic("please implement me")
}

//============================base64===========================
type base64Encryption struct {
}

func (base64Encryption base64Encryption) Encrypt(encodeStr []byte) ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(encodeStr)), nil
}

func (base64Encryption base64Encryption) Decrypt(decodeStr []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(string(decodeStr))
	if err != nil {
		return nil, err
	}
	return decodeBytes, nil
}
