package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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