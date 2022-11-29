package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteIpListRequest struct {
	IpgroupId string `json:"ipgroup_id"`

	Body *BatchDeleteIpListRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListRequest", string(data)}, " ")
}