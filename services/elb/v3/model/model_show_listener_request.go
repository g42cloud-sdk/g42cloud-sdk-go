package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowListenerRequest struct {
	ListenerId string `json:"listener_id"`
}

func (o ShowListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerRequest", string(data)}, " ")
}
