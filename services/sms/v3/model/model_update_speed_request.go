package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSpeedRequest struct {
	TaskId string `json:"task_id"`

	Body *SpeedLimit `json:"body,omitempty"`
}

func (o UpdateSpeedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpeedRequest struct{}"
	}

	return strings.Join([]string{"UpdateSpeedRequest", string(data)}, " ")
}
