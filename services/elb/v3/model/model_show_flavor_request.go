package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowFlavorRequest struct {
	FlavorId string `json:"flavor_id"`
}

func (o ShowFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorRequest struct{}"
	}

	return strings.Join([]string{"ShowFlavorRequest", string(data)}, " ")
}
