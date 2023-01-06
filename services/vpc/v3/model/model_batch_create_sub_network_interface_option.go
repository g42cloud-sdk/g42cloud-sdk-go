package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateSubNetworkInterfaceOption struct {
	VirsubnetId string `json:"virsubnet_id"`

	ParentId string `json:"parent_id"`

	SecurityGroups *[]string `json:"security_groups,omitempty"`

	Description *string `json:"description,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`
}

func (o BatchCreateSubNetworkInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceOption struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceOption", string(data)}, " ")
}
