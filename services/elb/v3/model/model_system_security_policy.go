package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SystemSecurityPolicy struct {
	Name string `json:"name"`

	Protocols string `json:"protocols"`

	Ciphers string `json:"ciphers"`

	ProjectId string `json:"project_id"`
}

func (o SystemSecurityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemSecurityPolicy struct{}"
	}

	return strings.Join([]string{"SystemSecurityPolicy", string(data)}, " ")
}
