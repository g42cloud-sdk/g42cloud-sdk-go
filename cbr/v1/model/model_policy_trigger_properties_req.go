package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTriggerPropertiesReq struct {
	Pattern []string `json:"pattern"`
}

func (o PolicyTriggerPropertiesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerPropertiesReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerPropertiesReq", string(data)}, " ")
}
