package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcPeeringRequestBody struct {
	Peering *CreateVpcPeeringOption `json:"peering"`
}

func (o CreateVpcPeeringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcPeeringRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcPeeringRequestBody", string(data)}, " ")
}
