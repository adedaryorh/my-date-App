package dtos

import "backend.app/common/constants"

type AreaOfInterest struct {
	AreasOfInterest []constants.AreaOfInterest `json:"areas_of_interest" validate:"gt=0,dive,is_enum"`
}
