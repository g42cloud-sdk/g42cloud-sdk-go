package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AllowDbPrivilegeResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AllowDbPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"AllowDbPrivilegeResponse", string(data)}, " ")
}
