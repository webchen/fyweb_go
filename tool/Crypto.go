package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	//"encoding/base64"
	"encoding/hex"
	"goweb/config"
)

// iv 长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
var iv = []byte(config.CRYPTOIV)


// Encrypt 加密函数
func Encrypt(plantText, key []byte) (string, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
	   return "", err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	//encrypted := base64.StdEncoding.EncodeToString(ciphertext)
	str := hex.EncodeToString(ciphertext)
	return str, nil
}

// PKCS7Padding 填充函数
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// Decrypt 解密函数
func Decrypt(ciphertext, key []byte) (string, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
	   return "", err
	}
	blockModel := cipher.NewCBCDecrypter(block, iv)
	ciphertext , _ = hex.DecodeString(string(ciphertext))
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return string(plantText), nil
}

// PKCS7UnPadding 反填充函数
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}