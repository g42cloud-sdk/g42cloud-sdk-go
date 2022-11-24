package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSubnetTagsRequest struct {
	SubnetId string `json:"subnet_id"`
}

func (o ShowSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowSubnetTagsRequest", string(data)}, " ")
}
