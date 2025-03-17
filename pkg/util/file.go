package util

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func IsExecl(filename string) bool {
	if strings.HasSuffix(filename, ".xlsx") || strings.HasSuffix(filename, ".xls") {
		return true
	}
	return false
}

func CreateTempFile(file *multipart.FileHeader) (*os.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	tempFile, err := os.CreateTemp("/tmp", "upload-*.xlsx")
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(tempFile, src); err != nil {
		return nil, err
	}

	return tempFile, nil
}
