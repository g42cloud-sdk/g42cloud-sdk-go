package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteIpAddressGroupForceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIpAddressGroupForceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpAddressGroupForceResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpAddressGroupForceResponse", string(data)}, " ")
}
