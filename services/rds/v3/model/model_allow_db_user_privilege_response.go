package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AllowDbUserPrivilegeResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AllowDbUserPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"AllowDbUserPrivilegeResponse", string(data)}, " ")
}
