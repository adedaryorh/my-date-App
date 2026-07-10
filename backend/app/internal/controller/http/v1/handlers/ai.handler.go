package handlers

import (
	"fmt"
	"net/http"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/configs"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/services/aiclient"
	"backend.app/pkg/response"
	"backend.app/pkg/logger"
	"github.com/gin-gonic/gin"
)

// AIHandler handles AI-related HTTP requests
type AIHandler struct {
	logger  *logger.Logger
	config  *configs.Config
	aiRyanClient *aiclient.AIServiceClient
}

// NewAIHandler creates a new AI handler instance
func NewAIHandler(logger *logger.Logger, config *configs.Config, aiRyanClient *aiclient.AIServiceClient) *AIHandler {
	return &AIHandler{
		logger:     logger,
		config:     config,
		aiRyanClient: aiRyanClient,
	}
}

// GetUserRecommendations godoc
// @Summary Get user recommendations
// @Description Get recommended users to follow based on AI
// @Tags AI
// @Accept json
// @Produce json
// @Param request body dtos.RecommendUsersRequest true "Recommendation request"
// @Success 200 {object} dtos.ResponseObject "Success"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /ai/users/recommendations [post]
func (h *AIHandler) GetUserRecommendations(c *gin.Context) {
	var input *dtos.RecommendUsersRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	user := c.MustGet("authUser").(models.User) // auth user

	// Override user_id from token for security
	input.UserID = user.ID.String()

	result := h.aiRyanClient.RecommendUsers(c.Request.Context(), input)
	if result != nil {
		responseObj := response.SuccessResponse("User recommendations retrieved successfully", result)
		c.JSON(http.StatusOK, responseObj)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to get user recommendations"), "Failed to get user recommendations")
	c.JSON(result.Code, result)
}

// ModerateCelebration godoc
// @Summary Moderate celebration content
// @Description Check celebration content for toxicity
// @Tags AI
// @Accept json
// @Produce json
// @Param request body dtos.ModerateCelebrationRequest true "Content to moderate"
// @Success 200 {object} dtos.ResponseObject "Success"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /ai/celebrations/moderate [post]
func (h *AIHandler) ModerateCelebration(c *gin.Context) {
	var input *dtos.ModerateCelebrationRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.aiRyanClient.ModerateCelebration(c.Request.Context(), input)
	if result != nil {
		responseObj := response.SuccessResponse("Content moderation completed", result)
		c.JSON(http.StatusOK, responseObj)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to moderate celebration"), "Failed to moderate celebration")
	c.JSON(result.Code, result)
}

// IndexCelebration godoc
// @Summary Index celebration for search
// @Description Index celebration content for semantic search
// @Tags AI
// @Accept json
// @Produce json
// @Param request body dtos.IndexCelebrationRequest true "Celebration to index"
// @Success 200 {object} dtos.ResponseObject "Success"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /ai/celebrations/index [post]
func (h *AIHandler) IndexCelebration(c *gin.Context) {
	var input *dtos.IndexCelebrationRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.aiRyanClient.IndexCelebration(c.Request.Context(), input)
	if result != nil {
		responseObj := response.SuccessResponse("Celebrity indexed successfully", result)
		c.JSON(http.StatusOK, responseObj)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to index celebration"), "Failed to index celebration")
	c.JSON(result.Code, result)
}

// SearchCelebrations godoc
// @Description Search celebrations using natural language queries
// @Tags AI
// @Accept json
// @Produce json
// @Param request body dtos.SearchCelebrationsRequest true "Search query"
// @Success 200 {object} dtos.ResponseObject "Success"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /ai/celebrations/search [post]
func (h *AIHandler) SearchCelebrations(c *gin.Context) {
	var input *dtos.SearchCelebrationsRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.aiRyanClient.SearchCelebrations(c.Request.Context(), input)
	if result != nil {
		responseObj := response.SuccessResponse("Celebrity search completed", result)
		c.JSON(http.StatusOK, responseObj)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to search celebrations"), "Failed to search celebrations")
	c.JSON(result.Code, result)
}

// LogInteraction godoc
// @Description Log user interactions (follow, skip, like, report) for AI feedback
// @Tags AI
// @Accept json
// @Produce json
// @Param request body dtos.LogInteractionRequest true "Interaction to log"
// @Success 200 {object} dtos.ResponseObject "Success"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /ai/interactions [post]
func (h *AIHandler) LogInteraction(c *gin.Context) {
	var input *dtos.LogInteractionRequest
	if err := c.Bind(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	result := h.aiRyanClient.LogInteraction(c.Request.Context(), input)
	if result != nil {
		responseObj := response.SuccessResponse("Interaction logged successfully", result)
		c.JSON(http.StatusOK, responseObj)
		return
	}

	// Handle error case
	result := response.ServerErrorResponse(fmt.Errorf("failed to log interaction"), "Failed to log interaction")
	c.JSON(result.Code, result)
}

// HealthCheck godoc
// @Description Check if AI service is healthy
// @Tags AI
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "Success"
// @Router /ai/health [get]
func (h *AIHandler) HealthCheck(c *gin.Context) {
	// For now, just return a simple OK response
	// In a real implementation, we might want to actually call the AI service health endpoint
	responseObj := response.SuccessResponse("AI service is healthy", map[string]string{
		"status":  "healthy",
		"service": "celebut-ai-service",
	})
	c.JSON(http.StatusOK, responseObj)
}