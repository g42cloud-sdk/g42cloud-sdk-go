package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTopicRequest struct {
	Body *CreateTopicRequestBody `json:"body,omitempty"`
}

func (o CreateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateTopicRequest", string(data)}, " ")
}
