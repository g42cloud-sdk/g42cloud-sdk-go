package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListNotifySmnTopicConfigRequest struct {
}

func (o ListNotifySmnTopicConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotifySmnTopicConfigRequest struct{}"
	}

	return strings.Join([]string{"ListNotifySmnTopicConfigRequest", string(data)}, " ")
}
