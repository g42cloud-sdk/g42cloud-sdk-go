package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Referer struct {
	RefererType int32 `json:"referer_type"`

	RefererList *string `json:"referer_list,omitempty"`

	IncludeEmpty *bool `json:"include_empty,omitempty"`
}

func (o Referer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Referer struct{}"
	}

	return strings.Join([]string{"Referer", string(data)}, " ")
}
