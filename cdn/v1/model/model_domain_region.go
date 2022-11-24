package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainRegion struct {
	DomainName *string `json:"domain_name,omitempty"`

	RegionIspDetails *[]map[string]interface{} `json:"region_isp_details,omitempty"`
}

func (o DomainRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainRegion struct{}"
	}

	return strings.Join([]string{"DomainRegion", string(data)}, " ")
}
