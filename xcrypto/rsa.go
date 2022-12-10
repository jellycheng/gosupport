package xcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 通过公钥进行rsa加密
func RsaEncrypt(data []byte, pubKeyFileName string) (error, []byte) {
	ret := make([]byte, 0, 0)
	f, err := os.Open(pubKeyFileName)
	if err != nil {
		return err, ret
	}
	fileInfo, err := f.Stat()
	if err != nil {
		return err, ret
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	buf := make([]byte, fileInfo.Size())
	_, _ = f.Read(buf)

	block, _ := pem.Decode(buf)

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err, ret
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		return err, ret
	}
	return nil, cipherText
}

// 通过私钥解密
func RsaDecrypt(data []byte, privateKeyFileName string) (error, []byte) {
	ret := make([]byte, 0, 0)
	f, err := os.Open(privateKeyFileName)
	if err != nil {
		return err, ret
	}
	fileInfo, err := f.Stat()
	if err != nil {
		return err, ret
	}
	buf := make([]byte, fileInfo.Size())
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	_, _ = f.Read(buf)
	block, _ := pem.Decode(buf)
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err, ret
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, data)
	if err != nil {
		return err, ret
	}
	return nil, plainText
}
