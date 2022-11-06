package validators

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"mime/multipart"
	"net/http"
)

// ValidateFileExtension is used to ensure the file is acceptable
func ValidateFileExtension(fh multipart.File, extensions []string) (string, error) {
	if fh == nil {
		return "", errors.New("header is nil")
	}

	buff := make([]byte, 512)
	if _, err := fh.Read(buff); err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}

	ext := http.DetectContentType(buff)
	if !slices.Contains(extensions, ext) {
		return ext, fmt.Errorf("invalid file extension: %s", ext)
	}

	return ext, nil
}
