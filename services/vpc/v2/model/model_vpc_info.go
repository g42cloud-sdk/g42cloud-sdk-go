package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VpcInfo struct {
	VpcId string `json:"vpc_id"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o VpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcInfo struct{}"
	}

	return strings.Join([]string{"VpcInfo", string(data)}, " ")
}
