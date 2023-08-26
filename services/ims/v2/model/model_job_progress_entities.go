package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type JobProgressEntities struct {
	ImageId *string `json:"image_id,omitempty"`

	CurrentTask *string `json:"current_task,omitempty"`

	ImageName *string `json:"image_name,omitempty"`

	ProcessPercent *float64 `json:"process_percent,omitempty"`

	SubJobId *string `json:"subJobId,omitempty"`
}

func (o JobProgressEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobProgressEntities struct{}"
	}

	return strings.Join([]string{"JobProgressEntities", string(data)}, " ")
}
