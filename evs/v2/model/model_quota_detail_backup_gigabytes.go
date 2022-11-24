package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailBackupGigabytes struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailBackupGigabytes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailBackupGigabytes struct{}"
	}

	return strings.Join([]string{"QuotaDetailBackupGigabytes", string(data)}, " ")
}
