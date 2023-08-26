package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpGroupRequest struct {
	IpgroupId string `json:"ipgroup_id"`

	Body *UpdateIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequest", string(data)}, " ")
}
