package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SubTaskAssociatedWithTask struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Progress *int32 `json:"progress,omitempty"`

	StartDate *int64 `json:"start_date,omitempty"`

	EndDate *int64 `json:"end_date,omitempty"`

	ProcessTrace *string `json:"process_trace,omitempty"`
}

func (o SubTaskAssociatedWithTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskAssociatedWithTask struct{}"
	}

	return strings.Join([]string{"SubTaskAssociatedWithTask", string(data)}, " ")
}
