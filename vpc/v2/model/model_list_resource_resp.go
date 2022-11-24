package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListResourceResp struct {
	ResouceDetail *interface{} `json:"resouce_detail"`

	ResourceId string `json:"resource_id"`

	ResourceName string `json:"resource_name"`

	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ListResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceResp struct{}"
	}

	return strings.Join([]string{"ListResourceResp", string(data)}, " ")
}
