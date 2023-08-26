package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateThumbnailsTaskRequest struct {
	Body *CreateThumbReq `json:"body,omitempty"`
}

func (o CreateThumbnailsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThumbnailsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateThumbnailsTaskRequest", string(data)}, " ")
}
