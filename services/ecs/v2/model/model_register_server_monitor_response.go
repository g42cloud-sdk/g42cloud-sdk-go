package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RegisterServerMonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterServerMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerMonitorResponse struct{}"
	}

	return strings.Join([]string{"RegisterServerMonitorResponse", string(data)}, " ")
}
