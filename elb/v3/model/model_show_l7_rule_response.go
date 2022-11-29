package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowL7RuleResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Rule           *L7Rule `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7RuleResponse struct{}"
	}

	return strings.Join([]string{"ShowL7RuleResponse", string(data)}, " ")
}
