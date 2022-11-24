package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailGigabytesSsd struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailGigabytesSsd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailGigabytesSsd struct{}"
	}

	return strings.Join([]string{"QuotaDetailGigabytesSsd", string(data)}, " ")
}
