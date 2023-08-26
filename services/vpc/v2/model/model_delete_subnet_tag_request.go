package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSubnetTagRequest struct {
	SubnetId string `json:"subnet_id"`

	Key string `json:"key"`
}

func (o DeleteSubnetTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubnetTagRequest", string(data)}, " ")
}
