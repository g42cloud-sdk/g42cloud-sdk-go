package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RejectVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o RejectVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"RejectVpcPeeringRequest", string(data)}, " ")
}
