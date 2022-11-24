package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpcPeeringOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateVpcPeeringOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcPeeringOption struct{}"
	}

	return strings.Join([]string{"UpdateVpcPeeringOption", string(data)}, " ")
}
