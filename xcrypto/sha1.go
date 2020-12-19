package xcrypto

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Sha1V1(str string) string  {
	h := sha1.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

