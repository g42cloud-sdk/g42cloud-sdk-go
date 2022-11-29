package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagResource struct {
	ResourceId string `json:"resource_id"`

	ResourceDetail []Vault `json:"resource_detail"`

	Tags []Tag `json:"tags"`

	ResourceName string `json:"resource_name"`

	SysTags []SysTag `json:"sys_tags"`
}

func (o TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResource struct{}"
	}

	return strings.Join([]string{"TagResource", string(data)}, " ")
}
