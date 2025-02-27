package core

import (
	"fmt"
	"mime/multipart"

	"backend.app/common/helpers"
	"backend.app/internal/models"
	"backend.app/internal/services/upload"
)

func (c *Core) uploadDocument(maxFileSize int64, kind, sourceId string, document *multipart.FileHeader, allowedKinds []upload.AttachmentKind) (*models.AWSObjectUrl, error) {
	file, err := helpers.MultipartHeaderToBytes(document)
	if err != nil {
		return nil, err
	}
	fileType := helpers.GetFileType(document.Filename, ".")
	// upload file to aws
	fileInput := upload.FileInput{
		Content: file,
		Name:    fmt.Sprintf("%s/%s.%s", kind, sourceId, fileType),
		Kind:    string(upload.AttachmentMap[fileType]),
		Size:    document.Size,
		URL:     fmt.Sprintf("%s/%s", kind, sourceId),
	}
	//upload file to AWS
	return c.UploadFileToAwsS3(fileInput, int64(maxFileSize), allowedKinds)

}
