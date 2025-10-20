package gosupport

import (
	"bytes"
	"compress/gzip"
	"io"
)

func IsGzipFile(data []byte) bool {
	return len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b
}

func GzipCompressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	_, err := gw.Write(data)
	if err != nil {
		return nil, err
	}
	gw.Close()
	// 必须关闭，确保所有数据写入缓冲区
	if err := gw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GzipDecompressData(compressed []byte) ([]byte, error) {
	buf := bytes.NewBuffer(compressed)
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, gr)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
