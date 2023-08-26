package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowServerRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRequest struct{}"
	}

	return strings.Join([]string{"ShowServerRequest", string(data)}, " ")
}
