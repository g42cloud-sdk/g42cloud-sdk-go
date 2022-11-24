package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDnsNameRequestBody struct {
	DnsType string `json:"dns_type"`
}

func (o CreateDnsNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDnsNameRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDnsNameRequestBody", string(data)}, " ")
}
