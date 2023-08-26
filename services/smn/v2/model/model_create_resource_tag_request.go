package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResourceTagRequest struct {
	ResourceType string `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequest", string(data)}, " ")
}
