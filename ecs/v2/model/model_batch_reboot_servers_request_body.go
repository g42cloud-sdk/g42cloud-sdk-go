package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchRebootServersRequestBody struct {
	Reboot *BatchRebootSeversOption `json:"reboot"`
}

func (o BatchRebootServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRebootServersRequestBody", string(data)}, " ")
}
