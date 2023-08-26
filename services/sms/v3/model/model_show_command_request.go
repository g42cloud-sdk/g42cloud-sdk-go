package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCommandRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowCommandRequest", string(data)}, " ")
}
