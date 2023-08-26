package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ServerSecurityGroup struct {
	Name string `json:"name"`

	Id string `json:"id"`
}

func (o ServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"ServerSecurityGroup", string(data)}, " ")
}
