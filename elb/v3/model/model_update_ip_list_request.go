package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIpListRequest struct {
	IpgroupId string `json:"ipgroup_id"`

	Body *UpdateIpListRequestBody `json:"body,omitempty"`
}

func (o UpdateIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpListRequest", string(data)}, " ")
}
