package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
