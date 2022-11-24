package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRestoreTimesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Date *string `json:"date,omitempty"`
}

func (o ListRestoreTimesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTimesRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreTimesRequest", string(data)}, " ")
}
