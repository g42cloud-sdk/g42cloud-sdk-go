package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteAnimatedGraphicsTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteAnimatedGraphicsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteAnimatedGraphicsTaskRequest", string(data)}, " ")
}
