package xcrypto

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//des ECB base64加密：xcrypto.EncryptDESECB2Base64("密码", "basesoft")
func EncryptDESECB2Base64(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

//des ECB hex加密
func EncryptDESECB2Hex(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}


//des ECB hex解密: src := xcrypto.DecryptDESECB4Hex(xcrypto.EncryptDESECB2Hex("123456", "basesoft"),"basesoft")
func DecryptDESECB4Hex(src, key string) string {
	data, err := hex.DecodeString(src)
	if err != nil {
		return ""
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return ""
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

//des ECB hex解密: src := xcrypto.DecryptDESECB4Base64(xcrypto.EncryptDESECB2Base64("hello'!@123", "basesoft"),"basesoft")
func DecryptDESECB4Base64(src, key string) string {
	data, err :=base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return ""
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}
