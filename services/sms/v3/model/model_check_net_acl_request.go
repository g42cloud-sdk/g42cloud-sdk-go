package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CheckNetAclRequest struct {
	TProjectId string `json:"t_project_id"`

	TNetworkId string `json:"t_network_id"`

	RegionId string `json:"region_id"`

	OsType string `json:"os_type"`
}

func (o CheckNetAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNetAclRequest struct{}"
	}

	return strings.Join([]string{"CheckNetAclRequest", string(data)}, " ")
}
