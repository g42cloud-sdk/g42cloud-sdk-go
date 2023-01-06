package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateIpGroupRequest struct {
	Body *CreateIpGroupRequestBody `json:"body,omitempty"`
}

func (o CreateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequest", string(data)}, " ")
}