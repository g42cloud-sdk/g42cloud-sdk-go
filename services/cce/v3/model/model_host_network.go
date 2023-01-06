package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type HostNetwork struct {
	Vpc string `json:"vpc"`

	Subnet string `json:"subnet"`

	SecurityGroup *string `json:"SecurityGroup,omitempty"`
}

func (o HostNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostNetwork struct{}"
	}

	return strings.Join([]string{"HostNetwork", string(data)}, " ")
}
