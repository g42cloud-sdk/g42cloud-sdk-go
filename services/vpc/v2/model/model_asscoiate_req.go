package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AsscoiateReq struct {
	Subnets *AssociateRouteTableAndSubnetReq `json:"subnets"`
}

func (o AsscoiateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsscoiateReq struct{}"
	}

	return strings.Join([]string{"AsscoiateReq", string(data)}, " ")
}
