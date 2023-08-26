package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteResetTracksTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteResetTracksTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskRequest", string(data)}, " ")
}
