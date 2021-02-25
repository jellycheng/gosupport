package xcrypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func Sha1V1(str string) string  {
	h := sha1.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256V1(str string) string  {
	h := sha256.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256ByByteV1(by []byte) string  {
	h := sha256.New()
	h.Write(by)
	return fmt.Sprintf("%x", h.Sum(nil))
}
