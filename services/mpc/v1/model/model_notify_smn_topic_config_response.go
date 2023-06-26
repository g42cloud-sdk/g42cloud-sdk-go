package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NotifySmnTopicConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NotifySmnTopicConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifySmnTopicConfigResponse struct{}"
	}

	return strings.Join([]string{"NotifySmnTopicConfigResponse", string(data)}, " ")
}
