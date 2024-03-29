package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// aes cbc加密
func AesCbcEncrypt(data, key []byte) []byte {
	// key长度为 16, 24 or 32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	data = Pkcs7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(data))
	blockMode.CryptBlocks(encrypted, data)
	return encrypted
}

// aes cbc解密
func AesCbcDecrypt(encrypted, key []byte) []byte {
	// key长度为 16, 24 or 32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = Pkcs7UnPadding(decrypted)
	return decrypted
}

// aes ecb加密
func AesEcbEncrypt(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(GetAesKey(key))
	length := (len(data) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, data)
	pad := byte(len(plain) - len(data))
	for i := len(data); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cipher.BlockSize(); bs <= len(data); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted
}

// aes ecb解密
func AesEcbDecrypt(encrypted, key []byte) []byte {
	cipher, _ := aes.NewCipher(GetAesKey(key))
	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}

// 加密，对应mysql功能： SELECT to_base64(AES_ENCRYPT("helloworld", "abc123key"));
func MysqlAesEncrypt(src, key string) string {
	bt := AesEcbEncrypt([]byte(src), []byte(key))
	return base64.StdEncoding.EncodeToString(bt)
}

// 解密，对应mysql功能： SELECT AES_DECRYPT(from_base64("irigEC0Ta0t8Bms/KO+Ocw=="), "abc123key")
func MysqlAesDecrypt(src, key string) string {
	if src, err := base64.StdEncoding.DecodeString(src); err == nil {
		bt := AesEcbDecrypt([]byte(src), []byte(key))
		return string(bt)
	} else {
		return ""
	}
}
