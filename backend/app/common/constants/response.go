package constants

type HttpStatus string

const (
	HttpStatusBadRequest       HttpStatus = "bad-request"
	HttpStatusServerError      HttpStatus = "server-error"
	HttpStatusSuccess          HttpStatus = "success"
	HttpStatusResourceNotFound HttpStatus = "resource-not-found"
	HttpStatusInvalidToken     HttpStatus = "invalid-token"
	HttpStatusTokenExpired     HttpStatus = "token-expired"
	HttpStatusTokenNotFound    HttpStatus = "token-not-found"
)
