package upload

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"backend.app/configs"
	"backend.app/internal/models"
	"backend.app/pkg/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

type (
	FileInput struct {
		Content []byte
		Name    string
		Kind    string
		URL     string
		Size    int64
		Type    *FileType
	}

	FileAttachment struct {
		Kind      string   `json:"kind"`
		URL       string   `json:"url"`
		Size      int64    `json:"size"`
		Extension string   `json:"extension"`
		Name      string   `json:"name,omitempty"`
		Type      FileType `json:"type,omitempty"`
	}

	AttachmentKind string
	FileType       string
)

const (
	AttachmentKindPDF       AttachmentKind = "application/pdf"
	AttachmentKindDoc       AttachmentKind = "application/doc"
	AttachmentKindImageJPEG AttachmentKind = "image/jpeg"
	AttachmentKindImageJPG  AttachmentKind = "image/jpg"
	AttachmentKindImagePNG  AttachmentKind = "image/png"
	AttachmentKindVideoMP4  AttachmentKind = "video/mp4"
	AttachmentKindVideoMOV  AttachmentKind = "video/mov"
)

var AttachmentKindMap = map[string]AttachmentKind{
	string(AttachmentKindPDF):       AttachmentKindPDF,
	string(AttachmentKindDoc):       AttachmentKindDoc,
	string(AttachmentKindImageJPEG): AttachmentKindImageJPEG,
	string(AttachmentKindImageJPG):  AttachmentKindImageJPG,
	string(AttachmentKindImagePNG):  AttachmentKindImagePNG,
	string(AttachmentKindVideoMP4):  AttachmentKindVideoMP4,
	string(AttachmentKindVideoMOV):  AttachmentKindVideoMOV,
}

var AttachmentMap = map[string]AttachmentKind{
	"pdf":  AttachmentKindPDF,
	"jpeg": AttachmentKindImageJPEG,
	"jpg":  AttachmentKindImageJPG,
	"png":  AttachmentKindImagePNG,
	"doc":  AttachmentKindDoc,
}

var MediaTypeMap = map[string]string{
	"mov":  "video",
	"mp4":  "video",
	"jpeg": "image",
	"jpg":  "image",
	"png":  "image",
}

// // SetAttachments sets Multiple FileAttachments
// func SetAttachments(attachments []FileAttachment) (*postgres.Jsonb, error) {
// 	j, err := json.Marshal(attachments)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &postgres.Jsonb{RawMessage: j}, nil
// }

// // SetAttachment sets Single FileAttachment
// func SetAttachment(attachment FileAttachment) (*postgres.Jsonb, error) {
// 	j, err := json.Marshal(attachment)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &postgres.Jsonb{RawMessage: j}, nil
// }

// // GetAttachments get the attachments in the FileAttachment format
// func GetAttachments(attach *postgres.Jsonb) ([]map[string]interface{}, error) {
// 	var attachments []map[string]interface{}
// 	b, err := json.Marshal(&attach)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(b, &attachments)
// 	return attachments, err
// }

type Uploader interface {
	ValidateFile(fileInput FileInput, maxFileSize int64, attachmentKinds []AttachmentKind) error
	UploadFile(fileInput FileInput) (*models.AWSObjectUrl, error)
}

type Upload struct {
	config *configs.Config
	log    *logger.Logger
}

func NewUpload(config *configs.Config, log *logger.Logger) *Uploader {
	upload := Uploader(&Upload{config: config})
	return &upload
}

func (u *Upload) s3FileSession() (*session.Session, error) {
	accessKeyID := u.config.AwsAccessKeyID
	secretAccessKey := u.config.AwsSecretAccessKey
	region := u.config.AwsRegion
	return session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKeyID,
				secretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
}

func (u *Upload) UploadFile(fileInput FileInput) (*models.AWSObjectUrl, error) {
	extension := filepath.Ext(fileInput.Name)

	// create a unique url for the file
	keyName := fmt.Sprintf("%s%s", fileInput.URL, extension)

	sess, err := u.s3FileSession()
	if err != nil {
		log.Err(err).Msgf("uploadFile:s3FileSession (%v)", err)
		return nil, err
	}

	output, err := s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Key:                  aws.String(keyName),
		Bucket:               aws.String(u.config.AwsS3Bucket),
		ACL:                  aws.String("public-read"),
		ContentType:          aws.String(http.DetectContentType(fileInput.Content)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
		Body:                 bytes.NewReader(fileInput.Content),
	})
	if err != nil {
		log.Err(err).Msgf("uploadFile:Upload (%v)", err)
		return nil, err
	}
	now := time.Now()
	return &models.AWSObjectUrl{
		KeyName:   keyName,
		Url:       output.Location,
		CreatedAt: &now,
	}, nil
}

// ValidateFile validates the file provided
func (u *Upload) ValidateFile(fileInput FileInput, maxFileSize int64, attachmentKinds []AttachmentKind) error {

	if fileInput.Size > maxFileSize {
		return fmt.Errorf("%s file should of size of %d, MB or less", fileInput.Name, maxFileSize)
	}
	allowed := false
	for _, kind := range attachmentKinds {
		if kind == AttachmentKindMap[fileInput.Kind] {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("file uploaded should be of the format(s): %s", aggregatedAttachmentKind(attachmentKinds))
	}

	return nil
}

// aggregatedAttachmentKind return joined, well formatted attachments
func aggregatedAttachmentKind(attachments []AttachmentKind) string {
	result := ""
	for _, attachment := range attachments {
		result += fmt.Sprintf("%s, ", attachment)
	}
	return result[:len(result)-2] // remove the comma and space for the last item
}
