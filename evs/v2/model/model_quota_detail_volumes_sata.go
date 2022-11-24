package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailVolumesSata struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailVolumesSata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailVolumesSata struct{}"
	}

	return strings.Join([]string{"QuotaDetailVolumesSata", string(data)}, " ")
}
