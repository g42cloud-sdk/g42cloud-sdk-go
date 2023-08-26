package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type LogObject struct {
	DomainName *string `json:"domain_name,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int64 `json:"size,omitempty"`

	Link *string `json:"link,omitempty"`
}

func (o LogObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogObject struct{}"
	}

	return strings.Join([]string{"LogObject", string(data)}, " ")
}
