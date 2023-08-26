package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RemuxRetryReq struct {
	TaskId *string `json:"task_id,omitempty"`
}

func (o RemuxRetryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemuxRetryReq struct{}"
	}

	return strings.Join([]string{"RemuxRetryReq", string(data)}, " ")
}
