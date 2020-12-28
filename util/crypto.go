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
type cryptoUtil struct {
	aes    aesEncryption
	base64 base64Encryption
	md5    md5Encryption
}

func NewCryptoUtil() *cryptoUtil {
	crypto := new(cryptoUtil)
	return crypto
}

///////////////////////////////补全方式的接口和实现////////////////////////
type pad interface {
	Padding(ciphertext []byte, blockSize int) []byte
	UnPadding(origData []byte) []byte
}

type PKCS5 struct {
}

func (P PKCS5) Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func (P PKCS5) UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

/////////////////////////////加密方式的接口和实现///////////////////////
type mode interface {
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
type aesEncryption struct {
	//key
	secretKey []byte
	//iv
	secretIv []byte
	//加密方式
	mode mode
	//补全方式
	pad pad
}

func (aesEncryption *aesEncryption) judgeOption() error {
	SecretKey := aesEncryption.secretKey
	SecretIv := aesEncryption.secretIv
	if len(SecretKey) != 16 && len(SecretKey) != 24 && len(SecretKey) != 32 {
		return errors.New("key length must be 16 or 24 or 32 byte")
	}
	if len(SecretKey) != len(SecretIv) {
		return errors.New("key and iv must be equal")
	}
	if aesEncryption.mode == nil || aesEncryption.pad == nil {
		return errors.New("mode and pad not empty")
	}
	return nil
}

func (aesEncryption *aesEncryption) SetOption(SecretKey []byte, SecretIv []byte, mode mode, pad pad) error {
	aesEncryption.secretKey = SecretKey
	aesEncryption.secretIv = SecretIv
	aesEncryption.mode = mode
	aesEncryption.pad = pad
	return aesEncryption.judgeOption()
}

func (aesEncryption *aesEncryption) Encrypt(encodeBytes []byte) ([]byte, error) {
	if err := aesEncryption.judgeOption(); err != nil {
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

func (aesEncryption *aesEncryption) Decrypt(decodeBytes []byte) ([]byte, error) {
	if err := aesEncryption.judgeOption(); err != nil {
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

//========================== md5 ===============================
type md5Encryption struct {
}

func (md5Encryption md5Encryption) Encrypt(encodeBytes []byte) ([]byte, error) {
	h := md5.New()
	h.Write(encodeBytes)
	return h.Sum(nil), nil
	//return hex.EncodeToString(h.Sum(nil)), nil
}

func (md5Encryption md5Encryption) Decrypt(decodeBytes []byte) ([]byte, error) {
	panic("please implement me")
}

//============================base64===========================
type base64Encryption struct {
}

func (base64Encryption base64Encryption) Encrypt(encodeBytes []byte) ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(encodeBytes)), nil
}

func (base64Encryption base64Encryption) Decrypt(decodeBytes []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(string(decodeBytes))
	if err != nil {
		return nil, err
	}
	return decodeBytes, nil
}
