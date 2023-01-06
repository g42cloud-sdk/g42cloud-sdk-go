package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmRulePoliciesResponse struct {
	Policies       *[]Policy `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateAlarmRulePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesResponse", string(data)}, " ")
}
