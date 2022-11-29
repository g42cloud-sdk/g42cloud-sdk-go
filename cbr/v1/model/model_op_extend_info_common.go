package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
