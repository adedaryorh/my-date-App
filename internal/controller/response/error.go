package response

import "fmt"

type ServiceErrorResponse struct {
	Message    string
	Err        error
	StatusCode int
}

func (e *ServiceErrorResponse) Error() string {
	return fmt.Sprintf("%v", e.Err)
}
