package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsReq struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o TagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsReq struct{}"
	}

	return strings.Join([]string{"TagsReq", string(data)}, " ")
}
