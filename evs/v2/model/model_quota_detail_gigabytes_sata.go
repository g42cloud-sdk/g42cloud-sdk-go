package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailGigabytesSata struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailGigabytesSata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailGigabytesSata struct{}"
	}

	return strings.Join([]string{"QuotaDetailGigabytesSata", string(data)}, " ")
}
