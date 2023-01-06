package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Statement struct {
	Sid string `json:"Sid"`

	Effect string `json:"Effect"`

	Principal *string `json:"Principal,omitempty"`

	NotPrincipal *string `json:"NotPrincipal,omitempty"`

	Action *string `json:"Action,omitempty"`

	NotAction *string `json:"NotAction,omitempty"`

	Resource *string `json:"Resource,omitempty"`

	NotResource *string `json:"NotResource,omitempty"`
}

func (o Statement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Statement struct{}"
	}

	return strings.Join([]string{"Statement", string(data)}, " ")
}
