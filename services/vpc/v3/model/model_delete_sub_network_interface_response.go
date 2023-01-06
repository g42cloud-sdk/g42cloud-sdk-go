package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSubNetworkInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubNetworkInterfaceResponse", string(data)}, " ")
}
