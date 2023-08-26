package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Login struct {
	SshKey *string `json:"sshKey,omitempty"`

	UserPassword *UserPassword `json:"userPassword,omitempty"`
}

func (o Login) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Login struct{}"
	}

	return strings.Join([]string{"Login", string(data)}, " ")
}
