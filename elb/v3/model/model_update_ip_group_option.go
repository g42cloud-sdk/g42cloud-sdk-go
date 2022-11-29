package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIpGroupOption struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	IpList *[]UpadateIpGroupIpOption `json:"ip_list,omitempty"`
}

func (o UpdateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupOption", string(data)}, " ")
}
