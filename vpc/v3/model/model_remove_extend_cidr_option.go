package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveExtendCidrOption struct {
	ExtendCidrs []string `json:"extend_cidrs"`
}

func (o RemoveExtendCidrOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveExtendCidrOption struct{}"
	}

	return strings.Join([]string{"RemoveExtendCidrOption", string(data)}, " ")
}
