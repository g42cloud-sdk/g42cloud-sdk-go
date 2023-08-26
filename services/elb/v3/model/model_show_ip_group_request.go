package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowIpGroupRequest struct {
	IpgroupId string `json:"ipgroup_id"`
}

func (o ShowIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRequest", string(data)}, " ")
}
