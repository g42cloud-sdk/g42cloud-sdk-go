package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsResp struct {
	Key *string `json:"key,omitempty"`

	Values *string `json:"values,omitempty"`
}

func (o TagsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResp struct{}"
	}

	return strings.Join([]string{"TagsResp", string(data)}, " ")
}
