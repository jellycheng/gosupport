package xcrypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 通过内容生成16位长度的aes key
func GetAesKey(key []byte) []byte {
	genKey := make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// 通过内容生成8位长度的des key
func GetDesKey(key []byte) []byte {
	genKey := make([]byte, 8)
	copy(genKey, key)
	for i := 8; i < len(key); {
		for j := 0; j < 8 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

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

func Pkcs7Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padText...)
}

func Pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
// 创建公钥、私钥
func CreateRsaKey(keySize int, priKeyFile, pubKeyFile string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}
	if priKeyFile == "" {
		priKeyFile = "rsa_private.pem"
	}
	f, err := os.Create(priKeyFile)
	if err != nil {
		return err
	}
	_ = pem.Encode(f, &block)
	_ = f.Close()

	// public key
	publicKey := privateKey.PublicKey
	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derpText,
	}
	if pubKeyFile == "" {
		pubKeyFile = "rsa_public.pem"
	}
	f, err = os.Create(pubKeyFile)
	if err != nil {
		return err
	}
	_ = pem.Encode(f, &block)
	_ = f.Close()
	return nil
}
