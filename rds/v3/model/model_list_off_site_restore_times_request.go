package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListOffSiteRestoreTimesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Date *string `json:"date,omitempty"`
}

func (o ListOffSiteRestoreTimesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteRestoreTimesRequest struct{}"
	}

	return strings.Join([]string{"ListOffSiteRestoreTimesRequest", string(data)}, " ")
}
