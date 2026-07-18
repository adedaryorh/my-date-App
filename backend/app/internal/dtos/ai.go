package dtos

type RecommendUsersRequest struct {
	UserID string `json:"user_id"`
	Limit  int    `json:"limit" validate:"omitempty,min=1,max=100"`
}

type ModerateCelebrationRequest struct {
	Text string `json:"text" validate:"required,min=1,max=5000"`
}

type IndexCelebrationRequest struct {
	CelebrationID string `json:"celebration_id" validate:"required,uuid"`
	Text          string `json:"text" validate:"required,min=1,max=5000"`
}

type SearchCelebrationsRequest struct {
	Query string `json:"query" validate:"required,min=1,max=500"`
	Limit int    `json:"limit" validate:"omitempty,min=1,max=100"`
}

type LogInteractionRequest struct {
	UserID   string                 `json:"user_id"`
	TargetID string                 `json:"target_id" validate:"required,uuid"`
	Action   string                 `json:"action" validate:"required,oneof=follow skip like report"`
	Metadata map[string]interface{} `json:"metadata"`
}
