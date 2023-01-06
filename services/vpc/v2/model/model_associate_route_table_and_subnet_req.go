package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AssociateRouteTableAndSubnetReq struct {
	Associate *[]string `json:"associate,omitempty"`

	Disassociate *[]string `json:"disassociate,omitempty"`
}

func (o AssociateRouteTableAndSubnetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouteTableAndSubnetReq struct{}"
	}

	return strings.Join([]string{"AssociateRouteTableAndSubnetReq", string(data)}, " ")
}
