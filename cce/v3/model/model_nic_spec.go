package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NicSpec struct {
	SubnetId *string `json:"subnetId,omitempty"`

	FixedIps *[]string `json:"fixedIps,omitempty"`

	IpBlock *string `json:"ipBlock,omitempty"`
}

func (o NicSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NicSpec struct{}"
	}

	return strings.Join([]string{"NicSpec", string(data)}, " ")
}
