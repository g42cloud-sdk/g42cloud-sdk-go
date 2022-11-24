package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetail struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
