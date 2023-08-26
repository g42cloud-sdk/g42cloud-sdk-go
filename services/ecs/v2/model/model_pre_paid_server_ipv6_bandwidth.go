package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PrePaidServerIpv6Bandwidth struct {
	Id *string `json:"id,omitempty"`
}

func (o PrePaidServerIpv6Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerIpv6Bandwidth struct{}"
	}

	return strings.Join([]string{"PrePaidServerIpv6Bandwidth", string(data)}, " ")
}
