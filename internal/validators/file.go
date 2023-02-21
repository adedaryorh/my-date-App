package validators

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"mime/multipart"
	"net/http"

	"github.com/vincent-petithory/dataurl"
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

// ValidateBase64Extension is used to ensure the base64 data file is acceptable
func ValidateBase64Extension(data string, extensions []string) (string, error) {
	if len(data) == 0 {
		return "", errors.New("data is empty")
	}

	dataURL, err := dataurl.DecodeString(data)
	if err != nil {
		return "", err
	}

	contentType := dataURL.MediaType.ContentType()
	if !slices.Contains(extensions, dataURL.MediaType.ContentType()) {
		return contentType, fmt.Errorf("invalid file extension: %s", contentType)
	}

	return contentType, nil
}
