package validators

import (
	"celebut-api/internal/dtos"
	"fmt"
)

const MaxPostMessageSize = 250

type ValidatePost interface {
	Validate(post dtos.NewPost) error
}

type PostValidator struct {
}

func NewPostValidator() *PostValidator {
	return &PostValidator{}
}

func (pv *PostValidator) Validate(p dtos.NewPost) error {
	if p.Message != nil && len(*p.Message) > MaxPostMessageSize {
		return fmt.Errorf("post number of characters can not be larger than %d", MaxPostMessageSize)
	}

	return nil
}
