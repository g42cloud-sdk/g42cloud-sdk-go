package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceGroupRequest struct {
	ContentType string `json:"Content-Type"`

	Body *CreateResourceGroupRequestBody `json:"body,omitempty"`
}

func (o CreateResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequest", string(data)}, " ")
}
