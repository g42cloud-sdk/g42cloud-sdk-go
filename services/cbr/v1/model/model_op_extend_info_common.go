package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OpExtendInfoCommon struct {
	Progress int32 `json:"progress"`

	RequestId string `json:"request_id"`

	TaskId *string `json:"task_id,omitempty"`
}

func (o OpExtendInfoCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoCommon struct{}"
	}

	return strings.Join([]string{"OpExtendInfoCommon", string(data)}, " ")
}
