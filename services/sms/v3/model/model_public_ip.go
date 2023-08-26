package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PublicIp struct {
	Type string `json:"type"`

	BandwidthSize int32 `json:"bandwidth_size"`

	BandwidthShareType *string `json:"bandwidth_share_type,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
