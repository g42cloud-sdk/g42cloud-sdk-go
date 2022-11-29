package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmRuleResourcesResponse struct {
	Resources *[][]Dimension `json:"resources,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRuleResourcesResponse", string(data)}, " ")
}
