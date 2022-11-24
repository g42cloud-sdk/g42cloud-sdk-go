package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RefererRsp struct {
	RefererType *int32 `json:"referer_type,omitempty"`

	RefererList *string `json:"referer_list,omitempty"`

	IncludeEmpty *bool `json:"include_empty,omitempty"`
}

func (o RefererRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefererRsp struct{}"
	}

	return strings.Join([]string{"RefererRsp", string(data)}, " ")
}
