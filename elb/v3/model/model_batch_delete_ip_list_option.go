package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteIpListOption struct {
	IpList *[]IpGroupIp `json:"ip_list,omitempty"`
}

func (o BatchDeleteIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListOption", string(data)}, " ")
}
