package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowsSpeedLimitsRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowsSpeedLimitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowsSpeedLimitsRequest struct{}"
	}

	return strings.Join([]string{"ShowsSpeedLimitsRequest", string(data)}, " ")
}
