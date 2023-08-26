package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RevokeSqlserverDbUserPrivilegeResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RevokeSqlserverDbUserPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeSqlserverDbUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"RevokeSqlserverDbUserPrivilegeResponse", string(data)}, " ")
}
