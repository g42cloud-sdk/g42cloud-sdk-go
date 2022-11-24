package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResetServersPasswordRequestBody struct {
	NewPassword string `json:"new_password"`

	DryRun *bool `json:"dry_run,omitempty"`

	Servers []ServerId `json:"servers"`
}

func (o BatchResetServersPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordRequestBody", string(data)}, " ")
}
