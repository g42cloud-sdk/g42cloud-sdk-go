package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReinstallServerWithoutCloudInitOption struct {
	Adminpass *string `json:"adminpass,omitempty"`

	Keyname *string `json:"keyname,omitempty"`

	Userid *string `json:"userid,omitempty"`

	Mode *string `json:"mode,omitempty"`
}

func (o ReinstallServerWithoutCloudInitOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithoutCloudInitOption struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithoutCloudInitOption", string(data)}, " ")
}
