package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePoliciesPriorityRequestBody struct {
	L7policies *[]BatchUpdatePriorityRequestBody `json:"l7policies,omitempty"`
}

func (o BatchUpdatePoliciesPriorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityRequestBody", string(data)}, " ")
}
