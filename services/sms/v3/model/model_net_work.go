package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NetWork struct {
	Name string `json:"name"`

	Ip string `json:"ip"`

	Netmask string `json:"netmask"`

	Gateway string `json:"gateway"`

	Mtu *int32 `json:"mtu,omitempty"`

	Mac string `json:"mac"`

	Id *string `json:"id,omitempty"`
}

func (o NetWork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetWork struct{}"
	}

	return strings.Join([]string{"NetWork", string(data)}, " ")
}
