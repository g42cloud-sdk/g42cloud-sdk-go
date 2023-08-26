package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListNotifyEventRequest struct {
}

func (o ListNotifyEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotifyEventRequest struct{}"
	}

	return strings.Join([]string{"ListNotifyEventRequest", string(data)}, " ")
}
