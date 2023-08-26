package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o ShowVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcRequest", string(data)}, " ")
}
