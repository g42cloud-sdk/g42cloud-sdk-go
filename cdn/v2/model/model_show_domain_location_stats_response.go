package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDomainLocationStatsResponse struct {
	GroupBy *string `json:"group_by,omitempty"`

	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainLocationStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainLocationStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainLocationStatsResponse", string(data)}, " ")
}
