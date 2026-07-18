package core

import (
	"context"
	"fmt"
	"strings"
	"time"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/services/aiclient"
	"backend.app/internal/services/upload"
	"backend.app/pkg/response"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// create celebration general and premium
func (c *Core) CreateCelebrationDto(ctx context.Context, user *models.User, data *dtos.CreateCelebrationDto) *dtos.ResponseObject {
	ownerId := user.ID.String()
	status := constants.CelebrationStatusActive
	if data.Owner == constants.CelebrationOwnerFriend {
		status = constants.CelebrationStatusPending
		if data.OwnerId != nil {
			ownerId = *data.OwnerId
		}
	}

	content := strings.TrimSpace(strings.Join([]string{data.Caption, data.Notes}, " "))
	var moderation *aiclient.ModerateCelebrationResponse
	if content != "" {
		result, moderationErr := c.aiClient.ModerateCelebration(ctx, &aiclient.ModerateCelebrationRequest{Text: content})
		if moderationErr != nil {
			c.log.Warn(fmt.Sprintf("AI moderation unavailable; continuing with normal creation: %v", moderationErr))
		} else {
			moderation = result
			if result.IsFlagged {
				status = constants.CelebrationStatusPending
			}
		}
	}
	// check expiry
	expiry := 24
	if data.ExpiresIn != nil {
		expiry = *data.ExpiresIn
	}
	CelebrationDate, _ := time.Parse(constants.DATE_LAYOUT, data.CelebrationDate)
	expiresAt := CelebrationDate.Add(time.Hour * time.Duration(expiry))
	friends := pq.StringArray(data.SelectedFriends)
	newCelebration := models.Celebration{
		Audience:        string(data.Audience),
		Owner:           string(data.Owner),
		OwnerIdentifier: "id",
		OwnerId:         ownerId,
		Status:          string(status),
		Frequency:       string(data.Frequency),
		ExpiresAt:       expiresAt,
		Long:            data.Longitude,
		Lat:             data.Latitude,
		CreatedBy:       user.ID,
		CelebrationKind: string(data.CelebrationKind),
		CelebrationType: string(constants.CelebrationTypeGeneral),
		CelebrationDate: CelebrationDate,
		Caption:         data.Caption,
		Notes:           data.Notes,
	}
	if len(data.SelectedFriends) > 0 {
		newCelebration.SelectedFriends = &friends
	}

	celebration, err := c.repo.CreateCelebration(ctx, &newCelebration)
	if err != nil {
		return response.ServerErrorResponse(err)
	}

	if moderation != nil && moderation.IsFlagged {
		_, err = c.postgres.Pool.Exec(ctx, `
			INSERT INTO moderation_queue
				(celebration_id, user_id, toxicity_score, flagged_content)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (celebration_id) DO NOTHING`,
			celebration.ID, user.ID, moderation.ToxicityScore, content,
		)
		if err != nil {
			return response.ServerErrorResponse(fmt.Errorf("queue flagged celebration: %w", err))
		}
	}

	if err = c.uploadCelebrationMedia(ctx, celebration.ID, ownerId, data); err != nil {
		return response.ServerErrorResponse(err)
	}

	if content != "" && (moderation == nil || !moderation.IsFlagged) {
		if _, indexErr := c.aiClient.IndexCelebration(ctx, &aiclient.IndexCelebrationRequest{
			CelebrationID: celebration.ID.String(), Text: content,
		}); indexErr != nil {
			c.log.Warn(fmt.Sprintf("Celebration created but semantic indexing failed: %v", indexErr))
		}
	}

	return response.SuccessResponse(constants.CelebrationCreatedSuccessfully, celebration)
}

func (c *Core) uploadCelebrationMedia(ctx context.Context, celebrationId uuid.UUID, ownerId string, data *dtos.CreateCelebrationDto) error {

	newMedia := models.Media{
		ObjectRef: "celebration",
		OwnerId:   ownerId,
		ObjectId:  celebrationId,
		CreatedAt: time.Now().UTC(),
	}
	if data.Media1 != nil {
		fileType := helpers.GetFileType(data.Media1.Filename, ".")
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindVideoMP4, upload.AttachmentKindVideoMOV}
		media1, err := c.uploadDocument(constants.MaxCelebrationFileSize, string(constants.DocumentKindCelebrationMedia), celebrationId.String(), data.Media1, attachmentKinds)
		if err != nil {
			return err
		}
		newMedia.MediaType = upload.MediaTypeMap[fileType]
		newMedia.MediaUrl = media1.Url
		_, err = c.repo.CreateMedia(ctx, &newMedia)
		if err != nil {
			return err
		}
	}
	if data.Media2 != nil {
		fileType := helpers.GetFileType(data.Media2.Filename, ".")
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindVideoMP4, upload.AttachmentKindVideoMOV}
		media2, err := c.uploadDocument(constants.MaxCelebrationFileSize, string(constants.DocumentKindCelebrationMedia), celebrationId.String(), data.Media2, attachmentKinds)
		if err != nil {
			return err
		}

		newMedia.MediaType = upload.MediaTypeMap[fileType]
		newMedia.MediaUrl = media2.Url
		_, err = c.repo.CreateMedia(ctx, &newMedia)
		if err != nil {
			return err
		}
	}

	if data.Media3 != nil {
		fileType := helpers.GetFileType(data.Media3.Filename, ".")
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindVideoMP4, upload.AttachmentKindVideoMOV}
		media3, err := c.uploadDocument(constants.MaxCelebrationFileSize, string(constants.DocumentKindCelebrationMedia), celebrationId.String(), data.Media3, attachmentKinds)
		if err != nil {
			return err
		}

		newMedia.MediaType = upload.MediaTypeMap[fileType]
		newMedia.MediaUrl = media3.Url
		_, err = c.repo.CreateMedia(ctx, &newMedia)
		if err != nil {
			return err
		}
	}

	if data.Media4 != nil {
		fileType := helpers.GetFileType(data.Media4.Filename, ".")
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindVideoMP4, upload.AttachmentKindVideoMOV}
		media4, err := c.uploadDocument(constants.MaxCelebrationFileSize, string(constants.DocumentKindCelebrationMedia), celebrationId.String(), data.Media4, attachmentKinds)
		if err != nil {
			return err
		}

		newMedia.MediaType = upload.MediaTypeMap[fileType]
		newMedia.MediaUrl = media4.Url
		_, err = c.repo.CreateMedia(ctx, &newMedia)
		if err != nil {
			return err
		}
	}

	if data.Media5 != nil {
		fileType := helpers.GetFileType(data.Media5.Filename, ".")
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindVideoMP4, upload.AttachmentKindVideoMOV}
		media5, err := c.uploadDocument(constants.MaxCelebrationFileSize, string(constants.DocumentKindCelebrationMedia), celebrationId.String(), data.Media5, attachmentKinds)
		if err != nil {
			return err
		}

		newMedia.MediaType = upload.MediaTypeMap[fileType]
		newMedia.MediaUrl = media5.Url
		_, err = c.repo.CreateMedia(ctx, &newMedia)
		if err != nil {
			return err
		}
	}

	if data.Media6 != nil {
		fileType := helpers.GetFileType(data.Media6.Filename, ".")
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindVideoMP4, upload.AttachmentKindVideoMOV}
		media6, err := c.uploadDocument(constants.MaxCelebrationFileSize, string(constants.DocumentKindCelebrationMedia), celebrationId.String(), data.Media6, attachmentKinds)
		if err != nil {
			return err
		}

		newMedia.MediaType = upload.MediaTypeMap[fileType]
		newMedia.MediaUrl = media6.Url
		_, err = c.repo.CreateMedia(ctx, &newMedia)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetAllCelebrations get self celebrations
func (c *Core) GetAllCelebrations(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject {
	result, err := c.repo.GetAllCelebrations(ctx, user, query)
	if err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.CelebrationsFetchedSuccessfully, result)
}

// GetSingleCelebration get single celebration
func (c *Core) GetSingleCelebration(ctx context.Context, celebrationId uuid.UUID) *dtos.ResponseObject {
	result, err := c.repo.GetCelebrationByField(ctx, helpers.Map{"id": celebrationId})
	if err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.CelebrationFetchedSuccessfully, result)
}

// DeclineCelebration decline celebration
func (c *Core) DeclineCelebration(ctx context.Context, celebrationId uuid.UUID, user *models.User) *dtos.ResponseObject {
	// delete media
	if err := c.repo.DeleteMedia(ctx, helpers.Map{"object_id": celebrationId, "owner_id": user.ID}); err != nil {
		response.ServerErrorResponse(err)
	}

	if err := c.repo.DeleteCelebration(ctx, helpers.Map{"id": celebrationId, "owner_id": user.ID}); err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.CelebrationDeclinedSuccessfully, nil)
}

// AcceptCelebration accept celebration
func (c *Core) AcceptCelebration(ctx context.Context, celebrationId uuid.UUID, user *models.User) *dtos.ResponseObject {
	if err := c.repo.UpdateCelebration(ctx, celebrationId, user, helpers.Map{"status": constants.CelebrationStatusActive}); err != nil {
		response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.CelebrationAcceptedSuccessfully, nil)
}
