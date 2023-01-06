package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmResponse", string(data)}, " ")
}
