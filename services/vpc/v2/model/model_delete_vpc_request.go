package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o DeleteVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRequest", string(data)}, " ")
}
