package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
