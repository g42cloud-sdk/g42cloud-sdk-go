package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaDetailPerVolumeGigabytes struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailPerVolumeGigabytes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailPerVolumeGigabytes struct{}"
	}

	return strings.Join([]string{"QuotaDetailPerVolumeGigabytes", string(data)}, " ")
}
