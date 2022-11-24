package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListJobInfoDetailRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	StartTime string `json:"start_time"`

	EndTime *string `json:"end_time,omitempty"`
}

func (o ListJobInfoDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInfoDetailRequest struct{}"
	}

	return strings.Join([]string{"ListJobInfoDetailRequest", string(data)}, " ")
}
