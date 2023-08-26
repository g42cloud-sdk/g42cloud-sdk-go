package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SetPostgresqlDbUserPwdResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetPostgresqlDbUserPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPostgresqlDbUserPwdResponse struct{}"
	}

	return strings.Join([]string{"SetPostgresqlDbUserPwdResponse", string(data)}, " ")
}
