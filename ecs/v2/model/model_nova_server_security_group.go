package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaServerSecurityGroup struct {
	Name *string `json:"name,omitempty"`
}

func (o NovaServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"NovaServerSecurityGroup", string(data)}, " ")
}
