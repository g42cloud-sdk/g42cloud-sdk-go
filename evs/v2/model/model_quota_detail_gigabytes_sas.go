package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailGigabytesSas struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailGigabytesSas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailGigabytesSas struct{}"
	}

	return strings.Join([]string{"QuotaDetailGigabytesSas", string(data)}, " ")
}
