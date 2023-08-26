package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeletePortRequest struct {
	PortId string `json:"port_id"`
}

func (o DeletePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortRequest struct{}"
	}

	return strings.Join([]string{"DeletePortRequest", string(data)}, " ")
}
