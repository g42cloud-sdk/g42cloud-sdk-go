package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UrlObject struct {
	Id *string `json:"id,omitempty"`

	Url *string `json:"url,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	TaskId *string `json:"task_id,omitempty"`

	TaskType *string `json:"task_type,omitempty"`
}

func (o UrlObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlObject struct{}"
	}

	return strings.Join([]string{"UrlObject", string(data)}, " ")
}
