package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcResourceTagRequest struct {
	VpcId string `json:"vpc_id"`

	Body *CreateVpcResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateVpcResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagRequest", string(data)}, " ")
}
