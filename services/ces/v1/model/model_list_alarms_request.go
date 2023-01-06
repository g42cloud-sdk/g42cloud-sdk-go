package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmsRequest struct {
	ContentType string `json:"Content-Type"`

	Limit *int32 `json:"limit,omitempty"`

	Order *string `json:"order,omitempty"`

	Start *string `json:"start,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
