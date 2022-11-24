package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkSubnet struct {
	SubnetID string `json:"subnetID"`
}

func (o NetworkSubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkSubnet struct{}"
	}

	return strings.Join([]string{"NetworkSubnet", string(data)}, " ")
}
