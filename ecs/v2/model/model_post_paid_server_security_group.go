package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostPaidServerSecurityGroup struct {
	Id *string `json:"id,omitempty"`
}

func (o PostPaidServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"PostPaidServerSecurityGroup", string(data)}, " ")
}
