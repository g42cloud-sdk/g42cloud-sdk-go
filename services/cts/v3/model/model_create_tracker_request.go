package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTrackerRequest struct {
	Body *CreateTrackerRequestBody `json:"body,omitempty"`
}

func (o CreateTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerRequest struct{}"
	}

	return strings.Join([]string{"CreateTrackerRequest", string(data)}, " ")
}
