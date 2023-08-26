package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type QuotaDetailVolumesSas struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailVolumesSas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailVolumesSas struct{}"
	}

	return strings.Join([]string{"QuotaDetailVolumesSas", string(data)}, " ")
}
