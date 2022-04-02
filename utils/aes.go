package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// AesCBCEncrypt aes cbc模式加密
func AesCBCPk5Encrypt(origData, key, iv []byte) ([]byte, error) {
	if len(origData) < 1 {
		return []byte(""), errors.New("crypted is empty")
	}
	if len(key) < 1 || len(iv) < 1 {
		return []byte(""), errors.New("key or iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	// 如果你不指定填充及加密模式的话，将会采用 CBC 模式和 PKCS5Padding 填充进行处理, 这里采用 pkcs5Padding
	origData = PKCS5Padding(origData, blockSize)

	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))

	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

// aes CBC 模式解密
func AesCBCPk5Decrypt(crypted, key, iv []byte) ([]byte, error) {
	if len(crypted) < 1 {
		return []byte(""), errors.New("crypted is empty")
	}
	if len(key) < 1 || len(iv) < 1 {
		return []byte(""), errors.New("key or iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))

	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	// origData = ZeroUnPadding(origData)

	return PKCS5UnPadding(origData)
}

// Aes cbc 加密, pkcs7 填充
func AesCBCPk7Encrypt(origData, key []byte, iv []byte) ([]byte, error) {
	if len(origData) < 1 {
		return []byte(""), errors.New("crypted is empty")
	}
	if len(key) < 1 {
		return []byte(""), errors.New("key is empty")
	}
	if len(iv) < 1 {
		return []byte(""), errors.New("iv is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// Aes cbc 解密, pkcs7 填充
func AesCBCPk7Decrypt(crypted, key []byte, iv []byte) ([]byte, error) {
	if len(crypted) < 1 {
		return []byte(""), errors.New("crypted is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 加入判断条件防止 panic
	blockSize := block.BlockSize()
	if len(key) < blockSize {
		return nil, errors.New("key too short")
	}
	if len(crypted)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	orgData := make([]byte, len(crypted))
	blockMode.CryptBlocks(orgData, crypted)

	return PKCS7UnPadding(orgData)
}

// AesCBCPk7DecryptHex hex解码, aes cbc模式解密
func AesCBCPk7EncryptBase64(orgBytes []byte, key, iv []byte) (string, error) {
	encryptBytes, err := AesCBCPk7Encrypt(orgBytes, key, iv)
	if err != nil {
		return "", err
	}

	// 进行base64 编码
	encryptStr := base64.StdEncoding.EncodeToString(encryptBytes)

	return encryptStr, nil
}

// AesCBCPk7DecryptHex hex解码, aes cbc模式解密
func AesCBCPk7EncryptBase64V2(orgBytes []byte, encodingAesKey string) (string, error) {
	var md5Str = Md5EncodeToString(encodingAesKey)
	encryptBytes, err := AesCBCPk7Encrypt(orgBytes, getAesKey(md5Str), getIv(md5Str))
	if err != nil {
		return "", err
	}

	// 进行base64 编码
	encryptStr := base64.StdEncoding.EncodeToString(encryptBytes)

	return encryptStr, nil
}

// AesCBCPk7Decrypt base64解码, aes cbc模式解密
func AesCBCPk7DecryptBase64(orgData string, key, iv []byte) ([]byte, error) {
	// 新进行 base64 解码
	orgByte, err := base64.StdEncoding.DecodeString(orgData)

	if err != nil {
		return []byte(""), err
	}

	decryptBytes, err := AesCBCPk7Decrypt(orgByte, key, iv)
	if err != nil {
		return []byte(""), err
	}

	return decryptBytes, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)

	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])

	if length-unpadding < 0 {
		return []byte(""), errors.New("padding length error")
	}

	return origData[:(length - unpadding)], nil
}

// PKCS7 填充
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7 填充
func PKCS7UnPadding(plantText []byte) ([]byte, error) {
	length := len(plantText)
	unpadding := int(plantText[length-1])

	if length-unpadding < 0 {
		return []byte(""), errors.New("padding length error")
	}

	return plantText[:(length - unpadding)], nil
}

func GetRealString(encodingAesKey string, data string) string {
	dataTmp, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	var md5Str = Md5EncodeToString(encodingAesKey)
	rs, err := AesCBCPk7Decrypt(dataTmp, getAesKey(md5Str), getIv(md5Str))
	if err != nil {
		panic(err)
	}

	return string(rs)
}

func GetEncodeString(encodingAesKey string, data string) string {
	rs, err := AesCBCPk7EncryptBase64V2([]byte(data), encodingAesKey)
	if err != nil {
		panic(err)
	}
	return rs
}

func getAesKey(key string) []byte {
	if len(key) != 32 {
		panic("error secret key")
	}
	return []byte(key[2:7] + key[11:15] + key[18:25])
}

func getIv(key string) []byte {
	if len(key) != 32 {
		panic("error secret key")
	}
	return []byte(key[4:9] + key[16:23] + key[25:29])
}
