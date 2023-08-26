package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteAddressGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddressGroupResponse", string(data)}, " ")
}
