package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailBackups struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailBackups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailBackups struct{}"
	}

	return strings.Join([]string{"QuotaDetailBackups", string(data)}, " ")
}
