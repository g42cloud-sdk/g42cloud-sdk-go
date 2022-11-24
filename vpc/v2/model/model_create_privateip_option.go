package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePrivateipOption struct {
	SubnetId string `json:"subnet_id"`

	IpAddress *string `json:"ip_address,omitempty"`
}

func (o CreatePrivateipOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateipOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateipOption", string(data)}, " ")
}
