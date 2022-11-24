package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchTagAddActionRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *BatchTagActionAddRequestBody `json:"body,omitempty"`
}

func (o BatchTagAddActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagAddActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagAddActionRequest", string(data)}, " ")
}
