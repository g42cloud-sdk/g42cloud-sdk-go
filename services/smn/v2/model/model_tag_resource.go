package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TagResource struct {
	ResourceId string `json:"resource_id"`

	ResourceDetail *ResourceDetail `json:"resource_detail"`

	Tags []ResourceTag `json:"tags"`

	ResourceName string `json:"resource_name"`
}

func (o TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResource struct{}"
	}

	return strings.Join([]string{"TagResource", string(data)}, " ")
}
