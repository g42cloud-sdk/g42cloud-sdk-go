package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Lts struct {
	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`

	LogGroupName *string `json:"log_group_name,omitempty"`

	LogTopicName *string `json:"log_topic_name,omitempty"`
}

func (o Lts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Lts struct{}"
	}

	return strings.Join([]string{"Lts", string(data)}, " ")
}
