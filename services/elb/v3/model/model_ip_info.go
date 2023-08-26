package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type IpInfo struct {
	Ip string `json:"ip"`

	Description string `json:"description"`
}

func (o IpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpInfo struct{}"
	}

	return strings.Join([]string{"IpInfo", string(data)}, " ")
}
