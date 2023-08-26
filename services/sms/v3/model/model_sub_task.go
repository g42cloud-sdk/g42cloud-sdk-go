package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SubTask struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Progress int32 `json:"progress"`

	StartDate *int64 `json:"start_date,omitempty"`

	EndDate *int64 `json:"end_date,omitempty"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	UserOp *string `json:"user_op,omitempty"`

	ProcessTrace *string `json:"process_trace,omitempty"`
}

func (o SubTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTask struct{}"
	}

	return strings.Join([]string{"SubTask", string(data)}, " ")
}
