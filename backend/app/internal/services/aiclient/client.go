package aiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sony/gobreaker"

	"backend.app/configs"
	"backend.app/pkg/logger"
)

// AIServiceClient handles communication with the Python AI service
type AIServiceClient struct {
	httpClient     *http.Client
	config         *configs.Config
	logger         *logger.Logger
	baseURL        string
	circuitBreaker *gobreaker.CircuitBreaker
}

// NewAIServiceClient creates a new AI service client
func NewAIServiceClient(config *configs.Config, logger *logger.Logger) *AIServiceClient {
	// Parse timeout from config (expects format like "5s")
	timeout, err := time.ParseDuration(config.AIServiceTimeout)
	if err != nil {
		// Default to 5 seconds if parsing fails
		timeout = 5 * time.Second
	}

	// Configure circuit breaker settings
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "AIService",
		MaxRequests: uint32(config.AIServiceCircuitBreakerMaxRequests),
		Interval:    time.Duration(config.AIServiceCircuitBreakerInterval) * time.Second,
		Timeout:     time.Duration(config.AIServiceCircuitBreakerTimeout) * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip if failure rate exceeds threshold
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= uint32(config.AIServiceCircuitBreakerThreshold) && failureRatio >= config.AIServiceFailureRateThreshold
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			logger.Info(fmt.Sprintf("Circuit breaker %s changed from %s to %s", name, from.String(), to.String()))
		},
	})

	return &AIServiceClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		config:         config,
		logger:         logger,
		baseURL:        config.AIServiceURL,
		circuitBreaker: cb,
	}
}

// ExecuteRequest performs an HTTP request with circuit breaker protection
func (c *AIServiceClient) ExecuteRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	// Execute the request through the circuit breaker
	result, err := c.circuitBreaker.Execute(func() (interface{}, error) {
		// Prepare request body if provided
		var jsonData []byte
		var requestErr error
		if body != nil {
			jsonData, requestErr = json.Marshal(body)
			if requestErr != nil {
				return nil, fmt.Errorf("failed to marshal request: %w", requestErr)
			}
		}

		// Create HTTP request
		url := fmt.Sprintf("%s%s", c.baseURL, path)
		var httpReq *http.Request
		if body != nil {
			httpReq, requestErr = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(jsonData))
		} else {
			httpReq, requestErr = http.NewRequestWithContext(ctx, method, url, nil)
		}
		if requestErr != nil {
			return nil, fmt.Errorf("failed to create request: %w", requestErr)
		}
		if body != nil {
			httpReq.Header.Set("Content-Type", "application/json")
		}

		// Execute HTTP request
		resp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, fmt.Errorf("failed to call AI service: %w", err)
		}
		defer resp.Body.Close()

		// Check status code
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("AI service returned non-OK status: %d", resp.StatusCode)
		}

		// Read response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		return responseBody, nil
	})

	if err != nil {
		return nil, err
	}

	return result.([]byte), nil
}

// EmbedProfileResponse represents the response from embedding a user profile
type EmbedProfileResponse struct {
	UserID    string    `json:"user_id"`
	Embedding []float64 `json:"embedding"`
	TextUsed  string    `json:"text_used,omitempty"`
}

type HealthCheckResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

type EmbedProfileRequest struct {
	UserID    string   `json:"user_id"`
	Bio       string   `json:"bio"`
	Interests []string `json:"interests"`
}

// RecommendUsersRequest represents the request for user recommendations
type RecommendUsersRequest struct {
	UserID          string `json:"user_id"`
	Limit           int    `json:"limit,omitempty"`
	ExcludeFollowed bool   `json:"exclude_followed,omitempty"`
	ExcludeBlocked  bool   `json:"exclude_blocked,omitempty"`
}

// RecommendUsersResponse represents the response from user recommendations
type RecommendUsersResponse struct {
	RecommendedUsers []map[string]interface{} `json:"recommended_users"`
}

// ModerateCelebrationRequest represents the request to moderate celebration content
type ModerateCelebrationRequest struct {
	Text string `json:"text"`
}

// ModerateCelebrationResponse represents the response from content moderation
type ModerateCelebrationResponse struct {
	ToxicityScore float64            `json:"toxicity_score"`
	IsFlagged     bool               `json:"is_flagged"`
	Categories    map[string]float64 `json:"categories,omitempty"`
}

// IndexCelebrationRequest represents the request to index a celebration for search
type IndexCelebrationRequest struct {
	CelebrationID string `json:"celebration_id"`
	Text          string `json:"text"`
}

// IndexCelebrationResponse represents the response from indexing a celebration
type IndexCelebrationResponse struct {
	CelebrationID string `json:"celebration_id"`
	Indexed       bool   `json:"indexed"`
}

// SearchCelebrationsRequest represents the request to search celebrations
type SearchCelebrationsRequest struct {
	Query string `json:"query"`
	Limit int    `json:"limit,omitempty"`
}

// SearchCelebrationsResponse represents the response from searching celebrations
type SearchCelebrationsResponse struct {
	Results []map[string]interface{} `json:"results"`
}

// LogInteractionRequest represents the request to log a user interaction
type LogInteractionRequest struct {
	UserID   string                 `json:"user_id"`
	TargetID string                 `json:"target_id"`
	Action   string                 `json:"action"` // follow, skip, like, report
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// LogInteractionResponse represents the response from logging an interaction
type LogInteractionResponse struct {
	Logged        bool   `json:"logged"`
	InteractionID string `json:"interaction_id,omitempty"`
}

// EmbedProfile sends a user profile to be embedded by the AI service
func (c *AIServiceClient) EmbedProfile(ctx context.Context, req *EmbedProfileRequest) (*EmbedProfileResponse, error) {
	c.logger.Info(fmt.Sprintf("Calling AI service to embed profile for user %s", req.UserID))

	responseBody, err := c.ExecuteRequest(ctx, http.MethodPost, "/ai/embed-profile", req)
	if err != nil {
		return nil, err
	}

	var response EmbedProfileResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func (c *AIServiceClient) HealthCheck(ctx context.Context) (*HealthCheckResponse, error) {
	responseBody, err := c.ExecuteRequest(ctx, http.MethodGet, "/ai/health", nil)
	if err != nil {
		return nil, err
	}
	var response HealthCheckResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to decode health response: %w", err)
	}
	return &response, nil
}

// RecommendUsers gets user recommendations from the AI service
func (c *AIServiceClient) RecommendUsers(ctx context.Context, req *RecommendUsersRequest) (*RecommendUsersResponse, error) {
	c.logger.Info(fmt.Sprintf("Calling AI service for user recommendations: %s", req.UserID))

	responseBody, err := c.ExecuteRequest(ctx, http.MethodPost, "/ai/recommend-users", req)
	if err != nil {
		return nil, err
	}

	var response RecommendUsersResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// ModerateCelebration checks celebration content for toxicity
func (c *AIServiceClient) ModerateCelebration(ctx context.Context, req *ModerateCelebrationRequest) (*ModerateCelebrationResponse, error) {
	c.logger.Info("Calling AI service to moderate celebration")

	responseBody, err := c.ExecuteRequest(ctx, http.MethodPost, "/ai/moderate-celebration", req)
	if err != nil {
		return nil, err
	}

	var response ModerateCelebrationResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// IndexCelebration indexes a celebration for semantic search
func (c *AIServiceClient) IndexCelebration(ctx context.Context, req *IndexCelebrationRequest) (*IndexCelebrationResponse, error) {
	c.logger.Info(fmt.Sprintf("Calling AI service to index celebration %s", req.CelebrationID))

	responseBody, err := c.ExecuteRequest(ctx, http.MethodPost, "/ai/index-celebration", req)
	if err != nil {
		return nil, err
	}

	var response IndexCelebrationResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// SearchCelebrations searches for celebrations using semantic similarity
func (c *AIServiceClient) SearchCelebrations(ctx context.Context, req *SearchCelebrationsRequest) (*SearchCelebrationsResponse, error) {
	c.logger.Info(fmt.Sprintf("Calling AI service to search celebrations: %s", req.Query))

	responseBody, err := c.ExecuteRequest(ctx, http.MethodPost, "/ai/search-celebrations", req)
	if err != nil {
		return nil, err
	}

	var response SearchCelebrationsResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// LogInteraction logs a user interaction for the feedback loop
func (c *AIServiceClient) LogInteraction(ctx context.Context, req *LogInteractionRequest) (*LogInteractionResponse, error) {
	c.logger.Info(fmt.Sprintf("Calling AI service to log %s interaction from %s to %s", req.Action, req.UserID, req.TargetID))

	responseBody, err := c.ExecuteRequest(ctx, http.MethodPost, "/ai/log-interaction", req)
	if err != nil {
		return nil, err
	}

	var response LogInteractionResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
