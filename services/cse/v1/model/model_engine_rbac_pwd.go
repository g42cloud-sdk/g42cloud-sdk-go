package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EngineRbacPwd struct {
	Pwd *string `json:"pwd,omitempty"`
}

func (o EngineRbacPwd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineRbacPwd struct{}"
	}

	return strings.Join([]string{"EngineRbacPwd", string(data)}, " ")
}
