package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostPaidServerEip struct {
	Iptype string `json:"iptype"`

	Bandwidth *PostPaidServerEipBandwidth `json:"bandwidth"`

	Extendparam *PostPaidServerEipExtendParam `json:"extendparam,omitempty"`
}

func (o PostPaidServerEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerEip struct{}"
	}

	return strings.Join([]string{"PostPaidServerEip", string(data)}, " ")
}
