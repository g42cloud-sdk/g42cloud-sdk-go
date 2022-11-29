package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleCondition struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o UpdateRuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleCondition struct{}"
	}

	return strings.Join([]string{"UpdateRuleCondition", string(data)}, " ")
}
