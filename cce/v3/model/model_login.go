package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
