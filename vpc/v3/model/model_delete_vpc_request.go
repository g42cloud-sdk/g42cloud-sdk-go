package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
