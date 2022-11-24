package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Auditlog struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int64 `json:"size,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`
}

func (o Auditlog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Auditlog struct{}"
	}

	return strings.Join([]string{"Auditlog", string(data)}, " ")
}
