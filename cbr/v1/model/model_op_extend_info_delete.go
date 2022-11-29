package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoDelete struct {
	BackupId string `json:"backup_id"`

	BackupName string `json:"backup_name"`
}

func (o OpExtendInfoDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoDelete struct{}"
	}

	return strings.Join([]string{"OpExtendInfoDelete", string(data)}, " ")
}
