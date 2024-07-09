package curl

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadFile struct {
	maxMemory int64
}

func (m *UploadFile) SetMaxMemory(maxVal int64) *UploadFile {
	m.maxMemory = maxVal
	return m
}

func (m UploadFile) FormFile(r *http.Request, name string) (*multipart.FileHeader, error) {
	if r.MultipartForm == nil && m.maxMemory > 0 {
		if err := r.ParseMultipartForm(m.maxMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}
	_ = f.Close()
	return fh, err
}

// 保存上传文件
func (m UploadFile) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func NewUploadFile() *UploadFile {
	ret := &UploadFile{
		maxMemory: 32 << 20, // 32 MB
	}
	return ret
}
