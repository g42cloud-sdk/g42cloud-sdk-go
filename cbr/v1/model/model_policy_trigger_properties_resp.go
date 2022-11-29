package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTriggerPropertiesResp struct {
	Pattern []string `json:"pattern"`

	StartTime *string `json:"start_time,omitempty"`
}

func (o PolicyTriggerPropertiesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerPropertiesResp struct{}"
	}

	return strings.Join([]string{"PolicyTriggerPropertiesResp", string(data)}, " ")
}
