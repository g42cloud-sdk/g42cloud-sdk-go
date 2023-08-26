package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcPeeringResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcPeeringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcPeeringResponse", string(data)}, " ")
}
