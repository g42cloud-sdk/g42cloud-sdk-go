package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSubNetworkInterfacesQuantityRequest struct {
}

func (o ShowSubNetworkInterfacesQuantityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfacesQuantityRequest struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfacesQuantityRequest", string(data)}, " ")
}
