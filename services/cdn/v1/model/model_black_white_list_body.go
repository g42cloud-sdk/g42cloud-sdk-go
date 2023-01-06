package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BlackWhiteListBody struct {
	Type int32 `json:"type"`

	IpList *[]string `json:"ip_list,omitempty"`
}

func (o BlackWhiteListBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteListBody struct{}"
	}

	return strings.Join([]string{"BlackWhiteListBody", string(data)}, " ")
}
