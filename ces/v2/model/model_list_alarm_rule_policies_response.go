package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmRulePoliciesResponse struct {
	Policies *[]Policy `json:"policies,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulePoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulePoliciesResponse", string(data)}, " ")
}
