package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UnlockTargetEcsRequest struct {
	TaskId string `json:"task_id"`
}

func (o UnlockTargetEcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockTargetEcsRequest struct{}"
	}

	return strings.Join([]string{"UnlockTargetEcsRequest", string(data)}, " ")
}
